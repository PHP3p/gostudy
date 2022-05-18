package v1

import (
	api "day02/api/v1"
	"github.com/gin-gonic/gin"
)

var (
	user = api.User
)

// 初始化v1版本路由

func InitRouter(r *gin.Engine) {

	rLogin(r)
	// 	用户路由组
	rUser(r)

	// 分类路由组

	// 文章路由组
}
