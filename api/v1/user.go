package v1

import (
	"day02/model"
	"day02/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type apiUser struct {
}

var User = new(apiUser)

func (u *apiUser) AddUser(c *gin.Context) {
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
func (u *apiUser) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 121, "data": 1, "msg": errmsg.GetErrMsg(200)})
}
func (u *apiUser) SaveUser(c *gin.Context) {

}
func (u *apiUser) DeletUser(c *gin.Context) {

}
