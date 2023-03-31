package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrdiio/go-jwt-auth/models"
	"github.com/mrdiio/go-jwt-auth/request"
	"github.com/mrdiio/go-jwt-auth/response"
)

type UserController struct {
	UserService models.UserService
}

func NewUserController(userService models.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (c *UserController) Create(ctx *gin.Context) {
	request := request.UserRequest{}
	if err := ctx.ShouldBind(&request); err != nil {
		response.ValidationError(ctx, "Validation error", err)
		return
	}

	user := models.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}

	if err := c.UserService.Create(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Response{
			Message: "Error creating user",
		})
		return
	}

	response.Success(ctx, "User created successfully", user.ToUserResponse())

}

func (c *UserController) FindAll(ctx *gin.Context) {
	users := c.UserService.FindAll()

	response.Success(ctx, "Users fetched successfully", users)
}
