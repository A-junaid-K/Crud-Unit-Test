package database

import (
	"log"

	"github.com/testing_ap/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	db := createDB()
	autoMigrate(db)
	return db
}

func createDB() *gorm.DB {
	DSN := "host=localhost user=ajk password=11012005 dbname=ajkdb port=5432"
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Ussankutti{})
}
