package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"wozaizhao.com/api/common"
	"wozaizhao.com/api/models"
)

type addPageReq struct {
	Title       string `form:"title" binding:"required"`       // 标题
	Author      string `form:"author" binding:"required"`      // 作者
	Abstract    string `form:"abstract" binding:"required"`    // 摘要
	Content     string `form:"content" binding:"required"`     // 内容
	ContentType uint   `form:"contentType" binding:"required"` // 内容类型
}

type addCategoriesReq struct {
	Type    string   `form:"type" binding:"required"`    // 类型
	Content []string `form:"content" binding:"required"` // 内容
}

// AddPage 添加页
func AddPage(c *gin.Context) {
	var addReq addPageReq
	if err := c.BindJSON(&addReq); err != nil {
		log.Error("Bind Error", err)
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

// AddCategories 添加分类
func AddCategories(c *gin.Context) {
	var addReq addCategoriesReq
	if err := c.BindJSON(&addReq); err != nil {
		log.Error("Bind Error", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	cateType, err := common.ParseInt(addReq.Type)
	if err != nil {
		log.Error("ParseFloat cateType", err)
	}
	for _, value := range addReq.Content {

		models.CreateCategory(uint(cateType), value)
	}
	log.Debug("addReq", addReq)
}

// GetPage 获取页
func GetPage(c *gin.Context) {
	id := c.Param("id")
	page, err := models.GetPageByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	cates, errGetCategory := models.GetCategoryByType(page.CategoryType)
	if errGetCategory != nil {
		log.Error(errGetCategory)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  page,
		"cates": cates,
	})
}
