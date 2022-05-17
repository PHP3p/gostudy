package v1

import (
	"day02/model"
	"day02/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	code := model.ExitsUser(user.Username)
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	if code == http.StatusOK {
		model.CreatUser(&user)
		code = errmsg.SUCCSE
	}
	c.JSON(http.StatusOK, gin.H{"code": code, "data": user, "msg": errmsg.GetErrMsg(code)})

}
func GetUser(c *gin.Context) {

}
func SaveUser(c *gin.Context) {

}
func DeletUser(c *gin.Context) {

}
