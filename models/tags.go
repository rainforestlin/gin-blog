package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
	State     int    `json:"state"`
}

func GetTags(pageNum, pageSize int, maps interface{}) (tags []Tag) {
	db.Model(&Tag{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Model(&Tag{}).Select("id").Where("name=?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool {
	db.Model(Tag{}).Create(&Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	})
	return true
}

func ExistTagById(id int) bool {
	var tag Tag
	db.Model(&Tag{}).Select("id").Where("id=?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func DeleteTag(id int) bool {
	db.Model(&Tag{}).Where("id=?", id).Delete(&Tag{})
	return true
}

func ModifyTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id=?", id).Update(data)
	return true
}

//models callback
//创建
func (tag *Tag) BeforeCreate(scope *gorm.Scope) (err error) {
	//tx.Model(tag).Update("created_on",time.Now().Unix())
	scope.SetColumn("created_on",time.Now().Unix())
	return
}

//func (tag *Tag) AfterCreate(scope *gorm.Scope) error {
//	return nil
//}

//更新和创建的回调方法
//func (tag *Tag) AfterSave(scope *gorm.Scope) error {
//	return nil
//}
//func (tag *Tag) BeforeSave(scope *gorm.Scope) error {
//	return nil
//}

//更新的回调方法
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("modified_on",time.Now().Unix())
	//tx.Model(tag).Update("modified_on",time.Now().Unix())
	return nil
}
//func (tag *Tag) AfterUpdate(scope gorm.Scope) error {
//	return nil
//}
//
////删除的回调方法
//func (tag *Tag) AfterDelete(scope gorm.Scope) error {
//	return nil
//}
//func (tag *Tag) BeforeDelete(scope gorm.Scope) error {
//	return nil
//}
//
////查询的回调方法
//func (tag *Tag) AfterFind(scope gorm.Scope) error {
//	return nil
//}
