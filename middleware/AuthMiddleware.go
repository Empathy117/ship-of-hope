package middleware

import (
	"github.com/empathy117/ship-of-hope/common"
	"github.com/empathy117/ship-of-hope/model"
	"github.com/empathy117/ship-of-hope/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// get authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			msgTokenFormatError := "权限不足"
			response.Response(ctx, http.StatusUnauthorized, 401, nil, msgTokenFormatError)
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, cliams, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			msgTokenInvalid := "权限错误: 1"
			response.Response(ctx, http.StatusUnauthorized, 401, nil, msgTokenInvalid)
			ctx.Abort()
			return
		}

		// user authentication passed, return user information
		userId := cliams.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// user
		if user.ID == 0 {
			msgUserInvalid := "权限错误: 2"
			response.Response(ctx, http.StatusUnauthorized, 401, nil, msgUserInvalid)
			ctx.Abort()
			return
		}

		// user valid
		ctx.Set("user", user)

		ctx.Next()
	}
}