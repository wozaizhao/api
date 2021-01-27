package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Page 页
type Page struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(40);NOT NULL;DEFAULT ''"`     // 标题
	Author      string `json:"author" gorm:"type:varchar(10);NOT NULL"`               // 作者
	ReadCount   int64  `json:"readCount" gorm:"type:int(8);NOT NULL;DEFAULT '0'"`     // 读数
	Abstract    string `json:"abstract" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"` // 摘要
	Type        uint   `json:"type" gorm:"type:smallint(1);NOT NULL;DEFAULT ''"`      // 类型
	Content     string `json:"content" gorm:"type:varchar(10);NOT NULL"`              // 内容
	ContentType uint   `json:"contentType" gorm:"type:varchar(10);NOT NULL"`          // 内容类型
}

// CreatePage 创建页
func CreatePage(title, author, abstract, content string, contentType uint) error {
	page := Page{Title: title, Author: author, Abstract: abstract, Content: content, ContentType: contentType, ReadCount: 0, Type: 1}
	result := DB.Create(&page)
	log.Debug("CreatePage RowsAffected", result.RowsAffected)
	return result.Error
}

// GetPageByID 获取页面
func GetPageByID(id string) (page Page, err error) {
	result := DB.First(&page, id)
	return page, result.Error
}
