package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"wozaizhao.com/api/models"
)

type addSchoolReq struct {
	Name       string `form:"name" binding:"required"`  // 学校名字
	CountyName string `json:"countyName"`               // 区
	Level      string `form:"level" binding:"required"` // 办学等级
}

type addSchoolDistrictReq struct {
	Address       string `form:"address" binding:"required"`  // 地址
	CountyName    string `json:"countyName"`                  // 区
	CommunityName string `form:"community"`                   // 小区名字
	SchoolID      uint   `form:"schoolID" binding:"required"` // 学校ID
}

// AddSchool 添加学校
func AddSchool(c *gin.Context) {
	var addReq addSchoolReq
	if err := c.BindJSON(&addReq); err != nil {
		log.Error("Bind Error", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	schoolID, err := models.CreateSchool(addReq.Name, addReq.CountyName, addReq.Level)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": schoolID})
}

// AddSchoolDistrict 添加招生地段
func AddSchoolDistrict(c *gin.Context) {
	var addReq addSchoolDistrictReq
	if err := c.BindJSON(&addReq); err != nil {
		log.Error("Bind Error", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := models.CreateSchoolDistrict(addReq.Address, addReq.CountyName, addReq.CommunityName, addReq.SchoolID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "AddSchoolDistrict successfully")
}
