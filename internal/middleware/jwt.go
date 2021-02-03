package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/julianlee107/blogWithGin/pkg/app"
	"github.com/julianlee107/blogWithGin/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			token string
			code  = errcode.Success
		)
		if s, exist := context.GetQuery("token"); exist {
			token = s
		} else {
			token = context.GetHeader("token")
		}
		if token == "" {
			code = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = errcode.UnauthorizedTokenExpire
				default:
					code = errcode.UnauthorizedTokenError
				}
			}
		}
		if code != errcode.Success {
			response := app.NewResponse(context)
			response.ToErrorResponse(code)
			context.Abort()
			return
		}
		context.Next()

	}
}
