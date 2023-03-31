package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrdiio/go-jwt-auth/config"
	"github.com/mrdiio/go-jwt-auth/router"
)

func main() {
	app := config.App()

	gin := gin.Default()

	router.Setup(app.DB, gin)

	gin.Run(":" + app.Env.Port)
}
