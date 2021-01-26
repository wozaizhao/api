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
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	err := models.CreatePlace(addReq.CountyName, addReq.Name, addReq.Address, addReq.ServiceTime, addReq.Phone)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "server error")
		return
	}

	c.JSON(http.StatusOK, "add successfully")

}
