package main

import (
	"gorm-playground-test/models"
	"gorm.io/gen"
	"gorm.io/gen/examples/dal"
)

func generate() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
		WithUnitTest:  true,
	})
	g.UseDB(dal.DB)

	g.ApplyBasic(models.Toy{}, models.Pet{})

	g.Execute()
}
