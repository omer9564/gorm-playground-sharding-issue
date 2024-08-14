package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm-playground-test/models"
	"gorm.io/gen/examples/dal"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

func initDB() *gorm.DB {
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
				suffixes := []string{"shard1", "shard2"}
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
	return db
}

func main() {
	initDB()
}
