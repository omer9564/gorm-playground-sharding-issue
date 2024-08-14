package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	ddal "gorm-playground-test/dal"
	"gorm-playground-test/models"
	"gorm.io/gen/examples/dal"
	"gorm.io/sharding"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	db := dal.ConnectDB("/tmp/sqlite.db")
	db.Use(sharding.Register(
		sharding.Config{
			DoubleWrite: true,
			ShardingKey: "name",
			ShardingAlgorithm: func(columnValue any) (suffix string, err error) {
				columnValueString, ok := columnValue.(string)
				if !ok {
					return columnValueString, fmt.Errorf("column value is not a string")
				}
				return columnValueString, nil
			},
			NumberOfShards: 2,
			ShardingSuffixs: func() []string {
				suffixes := []string{"name-shard-1", "name-shard-2"}
				return suffixes
			},
			PrimaryKeyGenerator: sharding.PKCustom,
			PrimaryKeyGeneratorFn: func(_ int64) int64 {
				ID, _ := uuid.NewUUID()
				return int64(ID.ID())
			},
		},
		models.Toy{},
	))
	dal.DB = db
	generate()
	if err := dal.DB.AutoMigrate(&models.Toy{}, &models.Pet{}); err != nil {
		panic(err)
	}
	ddal.SetDefault(db)
	result, err := ddal.Toy.WithContext(context.Background()).Where(ddal.Toy.Name.Eq("name-shard-1")).First()
	if err != nil {
		panic(err)
	}
	if result != nil {
		fmt.Println(result.Name)
	} else {
		fmt.Println("could not find result")
	}
}
