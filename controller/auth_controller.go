package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrdiio/go-jwt-auth/helper"
	"github.com/mrdiio/go-jwt-auth/models"
	"github.com/mrdiio/go-jwt-auth/request"
	"github.com/mrdiio/go-jwt-auth/response"
	"github.com/thanhpk/randstr"
)

type AuthController struct {
	UserService models.UserService
}

func NewAuthController(userService models.UserService) *AuthController {
	return &AuthController{
		UserService: userService,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	request := request.UserRequest{}
	if err := ctx.ShouldBind(&request); err != nil {
		response.ValidationError(ctx, "Validation error", err)
		return
	}

	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		response.Error(ctx, "Error hashing password", err)
		return
	}

	code := randstr.String(20)

	user := models.User{
		Name:             request.Name,
		Username:         request.Username,
		Password:         hashedPassword,
		Email:            request.Email,
		VerificationCode: code,
	}

	if err := c.UserService.Create(&user); err != nil {
		response.Error(ctx, "Error creating user", err)
		return
	}

	// Send email after user is created
	emailData := helper.EmailData{
		URL:       "http://localhost:8000/api/v1/auth/verify/" + code,
		FirstName: user.Name,
		Subject:   "Email Verification",
	}

	go helper.SendEmail(&user, &emailData)

	response.Success(ctx, "User created successfully", user.ToUserResponse())

}

func (c *AuthController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login",
	})
}

func (c *AuthController) Refresh(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Refresh",
	})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout",
	})
}
