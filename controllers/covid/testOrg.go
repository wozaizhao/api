package covid

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wozaizhao.com/api/models"
)

// GetCovidTestOrgs 获取js-sdk配置
func GetCovidTestOrgs(c *gin.Context) {
	orgs := models.CovidTestOrgList(0, 20)
	c.JSON(http.StatusOK, orgs)
}
