package jwt

import (
	"blogWithGin/pkg/errCode"
	"blogWithGin/pkg/logging"
	"blogWithGin/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		var data interface{}

		code = errCode.SUCCESS
		token := context.Query("token")
		if token == "" {
			token,err := context.Cookie("token")
			if err != nil {
				logging.Error(err)
			}
			if token == ""{
				code = errCode.INVALID_PARAMS
			}
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = errCode.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errCode.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != errCode.SUCCESS {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  errCode.GetMsg(code),
				"data": data,
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
