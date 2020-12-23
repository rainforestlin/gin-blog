package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/julianlee107/blogWithGin/conf"
	"log"
	"time"
)

//初始化model

var db *gorm.DB

type Model struct {
	gorm.Model
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	dbType = conf.DbType
	dbName = conf.DbName
	user = conf.DbUser
	password = conf.DbHost
	host = conf.DbHost
	tablePrefix = conf.DbTablePrefix
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	// 禁止默认表明的复数属形式
	db.SingularTable(true)
	db.LogMode(true)
	// 设置连接池中最大闲置连接数
	db.DB().SetMaxIdleConns(10)
	// 设置连接池中最大连接数
	db.DB().SetMaxOpenConns(100)
	// 设置连接的最大可复用时间
	db.DB().SetConnMaxLifetime(10*time.Second)
	// gorm与表自动同步
	db.AutoMigrate()
}

func CloseDB() {
	defer db.Close()
}

func GetDB() *gorm.DB  {
	return db
}