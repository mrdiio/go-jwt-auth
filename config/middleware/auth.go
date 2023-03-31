package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mrdiio/go-jwt-auth/config"
	"github.com/mrdiio/go-jwt-auth/helper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string
		cookie, err := ctx.Cookie("access_token")

		authHeader := ctx.GetHeader("Authorization")
		fields := strings.Fields(authHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		config := config.LoadEnv()
		sub, err := helper.ValidateToken(accessToken, config.AccessTokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		ctx.Set("x-user-id", sub["payload"])
		ctx.Next()

	}
}
