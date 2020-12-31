package main

import (
	"log"
	"net/http"

	"github.com/julianlee107/blogWithGin/global"
	"github.com/julianlee107/blogWithGin/internal/model"
	"github.com/julianlee107/blogWithGin/pkg/logger"
	"github.com/julianlee107/blogWithGin/pkg/setting"
	"github.com/julianlee107/blogWithGin/internal/routers"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	router := routers.NewRouter()

	service := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    10 * global.ServerSetting.ReadTimeout,
		WriteTimeout:   10 * global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	service.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err:%v", err)
	}
	err = setupDBEngine()

	if err != nil {
		log.Fatalf("init.setupDBEngine err:%v", err)
	}
}
