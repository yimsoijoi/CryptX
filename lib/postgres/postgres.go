package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s",
		"localhost", "postgres", "Godisgay!666", "5432")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("db connected")
	return db, nil
}
