package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Category 分类
type Category struct {
	gorm.Model
	Type    uint   `json:"type" gorm:"type:smallint(1);NOT NULL;DEFAULT ''"` // 类型
	Content string `json:"content" gorm:"type:varchar(25);NOT NULL"`         // 内容
}

// CreateCategory 创建分类
func CreateCategory(cateType uint, content string) error {
	cate := Category{Type: cateType, Content: content}
	result := DB.Create(&cate)
	log.Debug("CreateCategory RowsAffected", result.RowsAffected)
	return result.Error
}

// CreateCategories 批量创建分类
// func CreateCategories(cateType uint, contents string) error {
// cate := Category{Type: cateType, Content: content}
// result := DB.Create(&cate)
// log.Debug("CreateCategory RowsAffected", result.RowsAffected)
// return result.Error
// }

// GetCategoryByType 获取分类
func GetCategoryByType(cateType uint) (cates []Category, err error) {
	result := DB.Where("type = ?", cateType).Find(&cates)
	return cates, result.Error
}
