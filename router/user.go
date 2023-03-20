package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mrdiio/go-jwt-auth/controller"
	"github.com/mrdiio/go-jwt-auth/repository"
	"github.com/mrdiio/go-jwt-auth/service"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB, router *gin.RouterGroup) {

	repository := repository.NewUserRepo(db)
	service := service.NewUserService(repository)
	controller := controller.NewUserController(service)

	userRouter := router.Group("/user")

	userRouter.GET("/all", controller.FindAll)
	userRouter.POST("/create", controller.Create)

}
