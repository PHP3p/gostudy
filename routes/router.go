package routes

import (
	api "day02/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "day02/utils"

func InitRouter() {
	// 禁止日志颜色 处理window 控制台 【？？】 问题
	gin.DisableConsoleColor()
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "hello diy framework v1 api",
			})
		})
		v1.GET("user/:id", api.GetUser)
		v1.POST("user", api.AddUser)
		v1.PUT("user/:id", api.SaveUser)
		v1.DELETE("user/:id", api.SaveUser)
		// 	用户路由组

		// 分类路由组

		// 文章路由组

	}
	r.Run(utils.HttpPort)
}
