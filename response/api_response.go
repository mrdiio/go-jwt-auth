package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// func Error(ctx *gin.Context, message string, err error) {
// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{
// 		Message: message,
// 		Errors:  err.Error(),
// 	})
// }

func ValidationError(ctx *gin.Context, message string, err error) {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			errors = append(errors, e.Field()+" is required")
		case "min":
			errors = append(errors, e.Field()+" must be at least "+e.Param()+" characters long")
		case "max":
			errors = append(errors, e.Field()+" must be at most "+e.Param()+" characters long")
		case "email":
			errors = append(errors, "invalid email address")
		default:
			errors = append(errors, e.Error())
		}
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{
		Message: message,
		Errors:  errors,
	})

}

func Success(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Message: message,
		Data:    data,
	})
}
