package config

import (
	"log"

	"github.com/mrdiio/go-jwt-auth/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() *gorm.DB {
	var err error
	dsn := viper.GetString("DATABASE_URL")
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
	log.Println("Database connection closed")
	db.Close()
}
