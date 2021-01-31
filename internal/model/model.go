package model

import (
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julianlee107/blogWithGin/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

type Model struct {
	ID         uint32         `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	CreatedBy  string         `json:"created_by"`
	UpdatedAt  time.Time      `json:"updated_at"`
	ModifiedBy string         `json:"modified_by"`
	Deleted    gorm.DeletedAt `json:"deleted"`
	IsDel      uint8          `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	if databaseSetting.DBType == "mysql" {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			databaseSetting.Username,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBname,
			databaseSetting.Charset,
			databaseSetting.ParseTime,
		)
		// gorm v2版本的大改动
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			// 命名政策
			NamingStrategy: schema.NamingStrategy{
				// 表明前缀
				TablePrefix: databaseSetting.TablePrefix,
				// 使用单数表名
				SingularTable: false,
			},
		})
		if err != nil {
			return nil, err
		}
		myDB, _ := db.DB()
		myDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
		myDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
		return db, nil
	}
	return nil, errors.New("暂时未开放除MySQL外其他数据库连接")
}
