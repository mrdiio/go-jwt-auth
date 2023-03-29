package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, gin *gin.Engine) {
	baseRouter := gin.Group("/api")

	AuthRouter(db, baseRouter)

	UserRouter(db, baseRouter)
}
