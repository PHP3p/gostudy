package errmsg

const (
	SUCCSE = 200
	ERROR  = 500
	// i / code= 1888 .. .用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIHE    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
)

// l / code = 2088...文章模块的错误
// [ / code = 3888...分类模块的错误
var codemsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在!",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIHE:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
