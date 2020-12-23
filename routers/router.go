package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianlee107/blogWithGin/conf"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	gin.SetMode(conf.RunMode)

	r := gin.Default()
	router := r.Group("/")
	{
		router.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "pong",
			})
		})
		r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	r.Run(fmt.Sprintf(":%d", conf.HttpPort))
	return r
}
