package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"wozaizhao.com/api/services/qqmap"
)

// School 学校
type School struct {
	gorm.Model
	Name       string  `json:"schoolName" gorm:"type:varchar(40);NOT NULL"`       // 学校名称
	CountyName string  `json:"countyName" gorm:"type:varchar(5);DEFAULT ''"`      // 区
	Level      string  `json:"level" gorm:"type:varchar(40);NOT NULL"`            // 办学等级
	Type       uint    `json:"type" gorm:"type:smallint(6);NOT NULL;DEFAULT '0'"` // 学校类型 0 幼儿园 1 小学 2 中学
	Lng        float64 `json:"lng"`
	Lat        float64 `json:"lat"`
}

// CreateSchool 创建学校
func CreateSchool(name, countyName, level string) (uint, error) {
	res, err := qqmap.Geocoder(name, "上海")
	if err != nil {
		log.Error(err)
	}
	// place := Place{CountyName: countyName, Name: name, Address: address, ServiceTime: serviceTime, Lng: res.Result.Location.Lng, Lat: res.Result.Location.Lat, Phone: phone, Type: 1}
	school := School{Name: name, CountyName: countyName, Level: level, Lng: res.Result.Location.Lng, Lat: res.Result.Location.Lat}
	result := DB.Create(&school)
	log.Debug("CreateSchool RowsAffected", result.RowsAffected)
	return school.ID, result.Error
}
