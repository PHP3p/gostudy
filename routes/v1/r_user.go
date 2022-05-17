package v1

import (
	api "day02/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func rUser(r *gin.Engine) {

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
	}
}
