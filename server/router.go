package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"wozaizhao.com/api/controllers/app"
	"wozaizhao.com/api/controllers/user"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// allows all origins
	r.Use(cors.Default())

	// 公众号网页接口
	wx := r.Group("/wx")
	{
		// 获取js-sdk配置信息
		wx.GET("getConfig", user.WxGetConfig)
		// 添加页
		wx.POST("page", app.AddPage)
		// 获取页
		wx.GET("page/:id", app.GetPage)
		// 添加分类
		wx.POST("cate", app.AddCategories)
		// 添加地点
		wx.POST("place", app.AddPlace)
		// 添加学校
		wx.POST("school", app.AddSchool)
		// 添加招生地段
		wx.POST("schoolDistrict", app.AddSchoolDistrict)
		// 获取地点
		wx.GET("places", app.GetPlaces)
	}

	// 小程序相关接口
	// weapp := r.Group("/weapp")
	// {
	// 	weapp.POST("")
	// }

	//管理后台接口
	// admin := r.Group("/admin")
	// {
	// 	admin.POST("upload", qiniu.Upload)
	// }

	return r
}
