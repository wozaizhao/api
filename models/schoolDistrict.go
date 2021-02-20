package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"wozaizhao.com/api/services/qqmap"
)

// SchoolDistrict 招生地段
type SchoolDistrict struct {
	gorm.Model
	Address       string  `json:"address" gorm:"type:varchar(140);NOT NULL"`             // 地址
	CountyName    string  `json:"countyName" gorm:"type:varchar(5);DEFAULT ''"`          // 区
	SchoolID      uint    `json:"schoolId" gorm:"type:smallint(8);NOT NULL;DEFAULT '0'"` // 学校ID
	CommunityName string  `json:"communityName" gorm:"type:varchar(40);NOT NULL"`        // 小区名字
	Lng           float64 `json:"lng"`
	Lat           float64 `json:"lat"`
	Type          uint    `json:"type" gorm:"type:smallint(6);NOT NULL;DEFAULT '0'"`   // 学校类型 0 幼儿园 1 小学 2 中学
	Year          uint    `json:"year" gorm:"type:varchar(4);NOT NULL;DEFAULT '2020'"` // 学年
}

// CreateSchoolDistrict 创建招生地段
func CreateSchoolDistrict(address, countyName, communityName string, schoolID uint) error {
	addressFmt := strings.Replace(address, ":", "", -1)
	addressSplit := strings.Split(addressFmt, "（")
	res, err := qqmap.Geocoder("上海市浦东新区"+addressSplit[0], "上海")
	if err != nil {
		log.Error(err)
	}
	// schoolIDInt, _ := common.ParseInt(schoolID)
	district := SchoolDistrict{Address: address, CountyName: countyName, CommunityName: communityName, Lng: res.Result.Location.Lng, Lat: res.Result.Location.Lat, SchoolID: schoolID, Year: 2020}
	result := DB.Create(&district)
	log.Debug("CreateSchoolDistrict RowsAffected", result.RowsAffected)
	return result.Error
}
