package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func rUser(r *gin.Engine) {

	v1 := r.Group("api/v1/")
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "hello diy framework v1 api",
			})
		})

		v1.GET("user/:id", user.GetUser)
		v1.POST("user", user.AddUser)
		v1.PUT("user/:id", user.SaveUser)
		v1.DELETE("user/:id", user.SaveUser)
	}
}
