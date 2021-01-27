package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wozaizhao.com/api/models"
)

type addPageReq struct {
	Title       string `form:"title" binding:"required"`       // 标题
	Author      string `form:"author" binding:"required"`      // 作者
	Abstract    string `form:"abstract" binding:"required"`    // 摘要
	Content     string `form:"content" binding:"required"`     // 内容
	ContentType uint   `form:"contentType" binding:"required"` // 内容类型
}

// AddPage 添加页
func AddPage(c *gin.Context) {
	var addReq addPageReq
	if err := c.BindJSON(&addReq); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := models.CreatePage(addReq.Title, addReq.Author, addReq.Abstract, addReq.Content, addReq.ContentType)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "add successfully")

}

// GetPage 获取页
func GetPage(c *gin.Context) {
	id := c.Param("id")
	page, err := models.GetPageByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, page)
}
