package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julianlee107/blogWithGin/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//初始化model

var db *gorm.DB
var mysqlDB *sql.DB

type Model struct {
	gorm.Model
}

func init() {
	var (
		err                                               error
		dbName, user, password, host, tablePrefix string
	)
	dbName = conf.DbName
	user = conf.DbUser
	password = conf.DbHost
	host = conf.DbHost
	tablePrefix = conf.DbTablePrefix
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	user,
	password,
	host,
	dbName)
	db, err = gorm.Open(mysql.Open(dsn),&gorm.Config{
			NowFunc: func() time.Time {
			  return time.Now().Local()
			},NamingStrategy: schema.NamingStrategy{
				TablePrefix: tablePrefix,   // 表名前缀，`User` 的表名应该是 `blog_users`
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `blog_user`
			  },
		  })
	if err != nil {
		log.Println(err)
	}
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return tablePrefix + defaultTableName
	// }
	// // 禁止默认表明的复数属形式
	// db.SingularTable(true)
	// db.LogMode(true)
	// // 设置连接池中最大闲置连接数
	// db.DB().SetMaxIdleConns(10)
	// // 设置连接池中最大连接数
	// db.DB().SetMaxOpenConns(100)
	// // 设置连接的最大可复用时间
	// db.DB().SetConnMaxLifetime(10 * time.Second)
	mysqlDB,err = db.DB()
	mysqlDB.SetConnMaxIdleTime(24*time.Hour)
	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(100)
	// gorm与表自动同步
	db.AutoMigrate()
}

func CloseDB() {
	defer mysqlDB.Close()
}

func GetDB() *gorm.DB {
	return db
}
