package model

import (
	"errors"
	"fmt"

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
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedOn  uint32 `json:"created_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedOn uint32 `json:"modified_on"`
	ModifiedBy string `json:"modified_by"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
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
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   databaseSetting.TablePrefix,
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
