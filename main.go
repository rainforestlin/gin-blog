package main

import (
	"blogWithGin/pkg/setting"
	"blogWithGin/routers"
	"fmt"
	"net/http"
)

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
	server.ListenAndServe()
}
