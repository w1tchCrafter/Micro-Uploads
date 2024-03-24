package config

import (
	"log"
	"micro_uploads/internal/models"

	"gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
)

func ConnectDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatalf("Error starting database connection: %s\n", err.Error())
	}

	err = db.AutoMigrate(&models.FileModel{}, &models.UserModel{})
	return db, err
}
