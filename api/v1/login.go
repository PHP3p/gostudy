package v1

import (
	"day02/middleware"
	"day02/model"
	"day02/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var u model.User
	c.ShouldBindJSON(&u)
	code := model.CheckLogin(u.Username)
	// code := model.ExitsUser(u.Username)
	str := ""

	if code == errmsg.SUCCSE {
		str, _ = middleware.SetToken(u.Username)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": str,
		"msg":  errmsg.GetErrMsg(code),
	})
}
