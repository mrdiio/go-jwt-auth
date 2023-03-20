package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mrdiio/go-jwt-auth/config"
	"github.com/mrdiio/go-jwt-auth/router"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	//sync database

}

func main() {
	db := config.DatabaseConnection()

	defer config.CloseDB()

	gin := gin.Default()

	router.Setup(db, gin)

	gin.Run()
}
