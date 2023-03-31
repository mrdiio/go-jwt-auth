package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrdiio/go-jwt-auth/config"
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Response{
			Message: "Error hashing password",
		})
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Response{
			Message: "Error creating user",
		})

		return
	}

	// Send email after user is created
	emailData := helper.EmailData{
		URL:       os.Getenv("APP_URL") + "/api/v1/auth/verify/" + code,
		FirstName: user.Name,
		Subject:   "Email Verification",
	}

	go helper.SendEmail(user.Email, &emailData)

	response.Success(ctx, "User created successfully", user.ToUserResponse())

}

func (c *AuthController) Login(ctx *gin.Context) {
	var request request.LoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		response.ValidationError(ctx, "Validation error", err)
		return
	}

	res, err := c.UserService.VerifyCredentials(request.Email, request.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
			Message: "Invalid credentials",
		})
		return
	}

	env := config.LoadEnv()

	ctx.SetCookie("access_token", res.AccessToken, 3600, "/", env.AppUrl, false, true)
	ctx.SetCookie("refresh_token", res.RefreshToken, 3600, "/", env.AppUrl, false, true)

	response.Success(ctx, "User logged in successfully", res)
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
