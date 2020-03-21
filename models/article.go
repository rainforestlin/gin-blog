package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Model(&Article{}).Select("id").Where("id=?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return count
}

func GetArticles(pageNum, pageSize int, maps interface{}) (articls []Article) {
	db.Model(&Article{}).Preload("Tag").Where(maps).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articls)
	return
}

func GetArticle(id int) (article Article) {
	db.Model(&Article{}).Where("id=?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}
func AddArticle(data map[string]interface{})bool  {
	db.Model(&Article{}).Create(&Article{
		TagID:      data["tag_id"].(int),
		Title:      data["title"].(string),
		Desc:       data["desc"].(string),
		Content:    data["content"].(string),
		CreatedBy:  data["created_by"].(string),
		State:      data["state"].(int),
	})
	return true
}
func DeleteArticle(id int)bool  {
	db.Model(&Article{}).Where("id=?",id).Delete(&Article{})
	return true
}


func (article Article) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return
}

func (article Article) BeforeUpdate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return
}
