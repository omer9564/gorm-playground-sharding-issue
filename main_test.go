package main

import (
	"context"
	"fmt"
	ddal "gorm-playground-test/dal"
	"gorm-playground-test/models"
	"gorm.io/gen/examples/dal"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	db := initDB()
	if err := dal.DB.AutoMigrate(&models.Toy{}, &models.Pet{}); err != nil {
		panic(err)
	}
	ddal.SetDefault(db)
	result, err := ddal.Toy.WithContext(context.Background()).Where(ddal.Toy.Name.Eq("shard1"), ddal.Toy.OwnerType.Eq("dummy")).First()
	if err != nil {
		panic(err)
	}
	if result != nil {
		fmt.Println(result.Name)
	} else {
		fmt.Println("could not find result")
	}
}
