package middleware

import (
	"day02/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// var JwtKey = "wewwwwuuj"
var JwtKey = []byte("wewwwwuuj")

type Mycalims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// 生成token
func SetToken(name string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := Mycalims{
		UserName: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "name",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR_TOKEN_WRONG
	}
	return token, errmsg.SUCCSE
}

// 验证token  需要生成token的相同参数条件
func CheckToken(token string) (*Mycalims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &Mycalims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*Mycalims); setToken.Valid {

		return key, errmsg.SUCCSE
	}

	return nil, errmsg.ERROR
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenHeader := context.Request.Header.Get("Authorization")
		// tokenHeader := context.Header.
		code := errmsg.SUCCSE
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
		}
		// strings.HasPrefix("Bearer ")
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			context.Abort()
		}
		key, checkCode := CheckToken(checkToken[1])
		if checkCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			context.Abort()
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_WRONG
			context.Abort()
		}
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
		context.Set("username", key.UserName)
		context.Next()
	}
}
