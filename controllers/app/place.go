package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"wozaizhao.com/api/common"
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

// AddPlace 创建一个地址
func AddPlace(c *gin.Context) {
	var addReq addPlaceReq
	if err := c.Bind(&addReq); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Debug("addReq", addReq)

	lng, parseLngErr := common.ParseFloat(addReq.Lng)
	if parseLngErr != nil {
		log.Error("parseLngErr", parseLngErr)
	}
	lat, parseLatErr := common.ParseFloat(addReq.Lat)
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
	placeType := c.DefaultQuery("type", "1")
	fieldName := c.Query("fieldName")
	fieldValue := c.Query("fieldValue")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "15"))
	places, err := models.PlaceList(placeType, fieldName, fieldValue, pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, places)
}
