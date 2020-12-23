package conf

import (
	"fmt"

	"github.com/go-ini/ini"
)

var (
	RunMode string

	PageSize    int
	JWTSecret   string
	TokenExpire int
	Domain      string

	HttpPort     int
	ReadTimeout  int
	WriteTimeout int

	DbType        string
	DbUser        string
	DbPassword    string
	DbHost        string
	DbName        string
	DbTablePrefix string
)

func init() {
	mode, err := ini.Load("conf/mode.ini")
	if err != nil {
		fmt.Println("加载mode文件错误，请检查文件路径", err)
	}
	RunMode = mode.Section("mode").Key("RUN_MODE").MustString("debug")

	file, err := ini.Load(fmt.Sprintf("conf/%s/app.ini", RunMode))

	if err != nil {
		fmt.Println("加载app文件错误，请检查文件路径", err)
	}
	LoadApp(file)
	LoadServer(file)
	LoadDatabase(file)

}

func LoadApp(file *ini.File) {
	app := file.Section("app")
	PageSize = app.Key("PAGE_SIZE").MustInt(10)
	JWTSecret = app.Key("JWT_SECRET").MustString("")
	TokenExpire = app.Key("TOKEN_EXPIRE").MustInt(10)
	Domain = app.Key("DOMAIN").MustString("localhost")
}

func LoadServer(file *ini.File) {
	server := file.Section("server")

	HttpPort = server.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = server.Key("READ_TIMEOUT").MustInt(60)
	WriteTimeout = server.Key("WRITE_TIMEOUT").MustInt(60)

}

func LoadDatabase(file *ini.File) {
	database := file.Section("database")
	DbType = database.Key("TYPE").MustString("mysql")
	DbUser = database.Key("USER").MustString("root")
	DbPassword = database.Key("PASSWORD").MustString("")
	DbHost = database.Key("HOST").MustString("127.0.0.1:3306")
	DbName = database.Key("NAME").MustString("")
	DbTablePrefix = database.Key("TABLE_PREFIX").MustString("")
}
