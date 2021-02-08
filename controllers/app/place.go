package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"wozaizhao.com/api/models"
)

type addPlaceReq struct {
	CountyName    string `json:"countyName"`    // 区
	Name          string `json:"orgName"`       // 地点名称
	Address       string `json:"orgAddress"`    // 地址
	ServiceTime   string `json:"serviceTime"`   // 服务时间
	Phone         string `json:"phone"`         // 电话
	BusinessScope string `json:"businessScope"` // 业务范围
	Remark        string `json:"remark"`        // 备注
	Lng           string `json:"lng"`           // 经度
	Lat           string `json:"lat"`           // 纬度
}

func parseFloat(str string) (float64, error) {
	num, err := strconv.ParseFloat(str, 64)
	return num, err
}

// AddPlace 创建一个地址
func AddPlace(c *gin.Context) {
	var addReq addPlaceReq
	if err := c.Bind(&addReq); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Debug("addReq", addReq)

	lng, parseLngErr := parseFloat(addReq.Lng)
	if parseLngErr != nil {
		log.Error("parseLngErr", parseLngErr)
	}
	lat, parseLatErr := parseFloat(addReq.Lat)
	if parseLatErr != nil {
		log.Error("parseLatErr", parseLatErr)
	}

	err := models.CreatePlace(addReq.Name, addReq.Address, addReq.ServiceTime, addReq.Phone, addReq.BusinessScope, addReq.Remark, lng, lat)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "add successfully")

}

// GetPlaces 获取地点
func GetPlaces(c *gin.Context) {
	places, err := models.PlaceList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, places)
}
