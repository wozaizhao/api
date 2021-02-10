package models

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wozaizhao.com/api/common"
)

// DB 数据库
var DB *gorm.DB

// OpenDB 打开数据库
func OpenDB() {
	dsn := common.GetDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("OpenDB Error", err)
	} else {
		DB = db
		initTable()
	}

}

func initTable() {
	DB.AutoMigrate(&Place{}, &Page{}, &Category{})
}

// Paginate 分页
func Paginate(pageNum, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum == 0 {
			pageNum = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// Filter 按条件过滤
func Filter(placeType, fieldName, fieldValue string) func(db *gorm.DB) *gorm.DB {
	var whereStr string
	return func(db *gorm.DB) *gorm.DB {
		if fieldName != "" && fieldValue != "" {
			whereStr = fmt.Sprintf("type = ? AND %s LIKE ?", fieldName)
			return db.Where(whereStr, placeType, "%"+fieldValue+"%")
		} else {
			return db.Where("type = ?", placeType)
		}
	}
}
