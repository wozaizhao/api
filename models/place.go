package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Place 地点
type Place struct {
	gorm.Model
	CountyName    string  `json:"countyName" gorm:"type:varchar(5);DEFAULT ''"`              // 区
	Name          string  `json:"orgName" gorm:"type:varchar(40);NOT NULL"`                  // 地点名称
	Address       string  `json:"orgAddress" gorm:"type:varchar(100);NOT NULL;DEFAULT '50'"` // 地址
	ServiceTime   string  `json:"serviceTime" gorm:"type:varchar(60);NOT NULL;DEFAULT ''"`   // 服务时间
	Phone         string  `json:"phone" gorm:"type:varchar(60);NOT NULL;DEFAULT ''"`         // 电话
	BusinessScope string  `json:"businessScope" gorm:"type:varchar(255);DEFAULT ''"`         // 营业范围
	Remark        string  `json:"remark" gorm:"type:varchar(255);DEFAULT ''"`                // 备注
	Lng           float64 `json:"lng"`
	Lat           float64 `json:"lat"`
	Type          uint    `json:"type" gorm:"type:smallint(1);NOT NULL;DEFAULT ''"` // 类型
}

// CreatePlace 创建地点
func CreatePlace(name, address, serviceTime, phone, businessScope, remark string, lng, lat float64) error {
	// res, err := qqmap.Geocoder("上海市"+address, "上海")
	// if err != nil {
	// 	log.Error(err)
	// }
	// place := Place{CountyName: countyName, Name: name, Address: address, ServiceTime: serviceTime, Lng: res.Result.Location.Lng, Lat: res.Result.Location.Lat, Phone: phone, Type: 1}
	place := Place{Name: name, Address: address, ServiceTime: serviceTime, BusinessScope: businessScope, Remark: remark, Lng: lng, Lat: lat, Phone: phone, Type: 2}
	result := DB.Create(&place)
	log.Debug("CreatePlace RowsAffected", result.RowsAffected)
	return result.Error
}

// PlaceList 地点列表
func PlaceList(placeType string) (places []Place, err error) {
	result := DB.Where("type = ?", placeType).Find(&places)
	return places, result.Error
}
