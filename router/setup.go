package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, gin *gin.Engine) {
	baseRouter := gin.Group("/api")

	v1Router := baseRouter.Group("/v1")
	AuthRouter(db, v1Router)
	UserRouter(db, v1Router)
}
