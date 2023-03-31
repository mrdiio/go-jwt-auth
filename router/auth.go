package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mrdiio/go-jwt-auth/controller"
	"github.com/mrdiio/go-jwt-auth/repository"
	"github.com/mrdiio/go-jwt-auth/service"
	"gorm.io/gorm"
)

func AuthRouter(db *gorm.DB, router *gin.RouterGroup) {
	repository := repository.NewUserRepo(db)
	service := service.NewUserService(repository)
	controller := controller.NewAuthController(service)

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	authRouter := router.Group("/auth")
	authRouter.POST("/refresh", controller.Refresh)
	authRouter.POST("/logout", controller.Logout)
}
