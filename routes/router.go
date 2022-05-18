package routes

import (
	v1 "day02/routes/v1"
	"day02/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 禁止日志颜色 处理window 控制台 【？？】 问题
	gin.DisableConsoleColor()
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	// 		接口版本路由
	v1.InitRouter(r)

	r.Run(utils.HttpPort)
}
