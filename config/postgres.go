package config

import (
	"crud-go/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=docker dbname=crud-go port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o banco de dados PostgreSQL: %v", err)
	}

	err = db.AutoMigrate(&models.Item{})
	if err != nil {
		return nil, fmt.Errorf("erro ao rodar automigrate do banco de dados PostgreSQL: %v", err)
	}

	return db, nil
}
