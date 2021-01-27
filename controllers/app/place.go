package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wozaizhao.com/api/models"
)

type addPlaceReq struct {
	CountyName  string `json:"countyName"`  // 区
	Name        string `json:"orgName"`     // 地点名称
	Address     string `json:"orgAddress"`  // 地址
	ServiceTime string `json:"serviceTime"` // 服务时间
	Phone       string `json:"phone"`       // 电话
}

// AddPlace 创建一个地址
func AddPlace(c *gin.Context) {
	var addReq addPlaceReq
	if err := c.Bind(&addReq); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := models.CreatePlace(addReq.CountyName, addReq.Name, addReq.Address, addReq.ServiceTime, addReq.Phone)

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
