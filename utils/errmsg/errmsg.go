package errmsg

// code常量
const (
	SUCCESS = 200
	ERROR   = 500
	// code = 1000...用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// code = 2000...文章模块的错误
	ERROR_ARTICLE_NOT_EXIST = 2001
	// code = 3000...分类模块的错误
	ERROR_CATEGORY_USED      = 3001
	ERROR_CATEGORY_NOT_EXIST = 3002
)

// 利用状态码字典查询对应的错误提示信息
var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",
	// 用户
	ERROR_USERNAME_USED:    "用户已存在!",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NO_RIGHT:    "用户没有权限",
	// 文章
	ERROR_ARTICLE_NOT_EXIST: "文章不存在",

	// 分类
	ERROR_CATEGORY_USED:      "分类已存在",
	ERROR_CATEGORY_NOT_EXIST: "分类不存在",
}

// 获取msg的getter函数
func GetErrMsg(code int) string {
	return codeMsg[code]
}
