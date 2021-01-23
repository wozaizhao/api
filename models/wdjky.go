package models

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
)

// CovidTestOrg 检测机构
type CovidTestOrg struct {
	field.DefaultField `bson:",inline"`
	CountyName         string `bson:"countyName"`
	OrgName            string `bson:"orgName"`
	OrgAddress         string `bson:"orgAddress"`
	ServiceTime        string `bson:"serviceTime"`
	Phone              string `bson:"phone"`
}

// CreateCovidTestOrg 创建检测机构
func CreateCovidTestOrg() {

}

// UpdateCovidTestOrg 更新检测机构
func UpdateCovidTestOrg() {

}

// RemoveCovidTestOrg 删除检测机构
func RemoveCovidTestOrg() {

}

// CovidTestOrgList 检测机构列表
func CovidTestOrgList(skip int64, limit int64) []CovidTestOrg {
	batch := []CovidTestOrg{}
	wdjky.Find(ctx, bson.M{}).Sort("countyName").Skip(skip).Limit(limit).All(&batch)
	return batch
}

// GetCovidTestOrgByID 查找检测机构
func GetCovidTestOrgByID() {

}
