package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julianlee107/blogWithGin/conf"
	"github.com/julianlee107/blogWithGin/pkg/logging"
	"github.com/julianlee107/blogWithGin/routers"
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
	router := gin.Default()
	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})
	router = routers.InitRouter()
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", conf.HttpPort),
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       time.Duration(conf.ReadTimeout),
		ReadHeaderTimeout: 0,
		WriteTimeout:      time.Duration(conf.WriteTimeout),
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	//_ = server.ListenAndServe()
	//	优雅地重启
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()
	//	等待中断信号来，并在5秒钟之后终止服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	logging.Info("Shutdown Server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logging.Fatal("Shutdown Server", err)
	}
	logging.Info("Server exiting")
}
