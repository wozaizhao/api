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

		wx.GET("getPage", app.GetPage)

		// 添加地点
		wx.POST("addPlace", app.AddPlace)

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
