package v1

import (
	v1 "day02/api/v1"
	"github.com/gin-gonic/gin"
)

func rLogin(r *gin.Engine) {
	r.POST("/login", v1.Login)
}
