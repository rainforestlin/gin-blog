package main

import (
	"blogWithGin/pkg/setting"
	"blogWithGin/routers"
	"fmt"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description 用gin框架搭建一个博客api.
// @contact.name Julian Lee
// @contact.email julianlee107@hotmail.com
// @host localhost:8000
// @BasePath /
// @securitydefinitions.oauth2.application JWT
// @scope.write Grants write access
// @tokenUrl 127.0.0.1:8000/auth
func main() {
	//router := gin.Default()
	//router.GET("/test", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "test",
	//	})
	//})
	router := routers.InitRouter()
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       setting.ReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      setting.WriteTimeout,
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	_ = server.ListenAndServe()
}
