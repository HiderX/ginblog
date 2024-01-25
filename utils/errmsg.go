package utils

const (
	SUCCESS = 200
	ERROR   = 500
	// code = 1000... 用户模块的错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST    = 1004
	ERROR_TOKEN_RUNTIME  = 1005
	ERROR_TOKEN_WRONG    = 1006
	ERROR_TOKEN_TYPE     = 1007

	// code = 2000... 分类模块的错误
	ERROR_CATEGORY_USED      = 2001
	ERROR_CATEGORY_NOT_EXIST = 2002

	// code = 3000... 文章模块的错误
	ERROR_ART_NOT_EXIST = 3001
)

var CodeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",
	// code = 1000... 用户模块的错误
	ERROR_USERNAME_USED:  "用户名已存在！",
	ERROR_PASSWORD_WRONG: "密码错误！",
	ERROR_USER_NOT_EXIST: "用户不存在！",
	ERROR_TOKEN_EXIST:    "TOKEN不存在！",
	ERROR_TOKEN_RUNTIME:  "TOKEN已过期！",
	ERROR_TOKEN_WRONG:    "TOKEN不正确！",
	ERROR_TOKEN_TYPE:     "TOKEN格式错误！",
	// code = 2000... 分类模块的错误
	ERROR_CATEGORY_USED:      "分类已存在！",
	ERROR_CATEGORY_NOT_EXIST: "分类不存在！",

	// code = 3000... 文章模块的错误
	ERROR_ART_NOT_EXIST: "文章不存在！",
}

func GetErrMsg(code int) string {
	return CodeMsg[code]
}
