package db

import (
	"log"
	"os"

	"github.com/gitarchived/api/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Create() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("PG_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	// Check if repository and host table exists
	if !db.Migrator().HasTable(&data.Repository{}) {
		db.Migrator().CreateTable(&data.Repository{})
	}

	if !db.Migrator().HasTable(&data.Host{}) {
		db.Migrator().CreateTable(&data.Host{})
	}

	return db
}
