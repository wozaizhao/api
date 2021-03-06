package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Page 页
type Page struct {
	gorm.Model
	Title         string `json:"title" gorm:"type:varchar(40);NOT NULL;DEFAULT ''"`     // 标题
	Author        string `json:"author" gorm:"type:varchar(10);NOT NULL"`               // 作者
	ReadCount     int64  `json:"readCount" gorm:"type:int(8);NOT NULL;DEFAULT '100'"`   // 读数
	Abstract      string `json:"abstract" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"` // 摘要
	Type          uint   `json:"type" gorm:"type:smallint(6);NOT NULL;DEFAULT ''"`      // 类型
	Content       string `json:"content" gorm:"type:varchar(10);NOT NULL"`              // 内容
	ContentType   uint   `json:"contentType" gorm:"type:smallint(6);NOT NULL"`          // 内容类型
	CategoryField string `json:"categoryField" gorm:"type:varchar(20)"`                 // 分类字段名
	CategoryType  uint   `json:"categoryType" gorm:"type:smallint(6)"`                  // 分类字段
}

// CreatePage 创建页
func CreatePage(title, author, abstract, content string, contentType uint) error {
	page := Page{Title: title, Author: author, Abstract: abstract, Content: content, ContentType: contentType, ReadCount: 0, Type: 1}
	result := DB.Create(&page)
	log.Debug("CreatePage RowsAffected", result.RowsAffected)
	return result.Error
}

// GetPageByID 获取页面
func GetPageByID(id uint) (page Page, err error) {
	result := DB.First(&page, id)
	page.ReadCount = page.ReadCount + 1
	DB.Save(&page)
	return page, result.Error
}
