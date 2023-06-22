package db

import (
	"log"

	"github.com/paraswaykole/transactionservice/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(postgres.Open(config.GetPGDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to db: %s", err)
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS ltree")
}

func Get() *gorm.DB {
	return db
}
