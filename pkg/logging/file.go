package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtimes\\logs\\"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}
func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permisson: %v",err)
	}
	handle,err := os.OpenFile(filePath,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		log.Fatalf("Fail to openFile:%v",err)
	}
	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	file,err := os.Create(dir+"\\"+getLogFileFullPath())
	defer file.Close()
	if err != nil {
		panic(err)
	}
}
