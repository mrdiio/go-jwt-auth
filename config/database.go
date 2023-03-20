package config

import (
	"os"

	"github.com/mrdiio/go-jwt-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() *gorm.DB {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&models.User{}, &models.Article{})

	return DB

}

func CloseDB() {
	db, err := DB.DB()
	if err != nil {
		panic("Failed to close database!")
	}
	db.Close()
}
