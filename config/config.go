package config

import (
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error

	db, err = InitializePostgres()
	if err != nil {
		return fmt.Errorf("erro ao inicializar o PostgreSQL: %v", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
