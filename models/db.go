package models

import (
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
