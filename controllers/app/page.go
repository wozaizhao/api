package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"wozaizhao.com/api/services/qqmap"
)

// GetPage 获取页
func GetPage(c *gin.Context) {
	// models.CreatePlace()
	res, err := qqmap.Geocoder("闵行区华漕社区卫生服务中心", "上海")
	log.Debug(res)
	if err != nil {
		log.Error(err)
	}
	c.JSON(http.StatusOK, "")
}
