package model

// 全局错误相关
const (
	Success  = 200
	Error    = 500
	ErrInner = 600

	// ERROR_USERNAME_USED code= 1000... 用户模块的错误
	ErrUserNameUsed      = 1001
	ErrPassWordWrong     = 1002
	ErrUserNotExists     = 1003
	ErrTokenExists       = 1004
	ErrTokenRuntime      = 1005
	ErrTokenWrong        = 1006
	ErrTokenTypeWrong    = 1007
	ErrUserNoRight       = 1008
	ErrUserNameOrPwWrong = 1009

	// code= 2000... 文章模块的错误
	ErrArtNotExists = 2001
	ErrArtExists    = 2002

	// code= 3000... 分类模块的错误
	ErrCateNameUsed  = 3001
	ErrCateNotExists = 3002

	// 查询主页不在
	ErrProfileNotExists = 4001
)

var Err2Msg = map[int]string{

	Success:              "OK",
	Error:                "FAIL",
	ErrUserNameUsed:      "用户名已存在！",
	ErrPassWordWrong:     "密码错误",
	ErrUserNameOrPwWrong: "用户名或密码错误",
	ErrUserNotExists:     "用户不存在",
	ErrTokenExists:       "TOKEN不存在,请重新登陆",
	ErrTokenRuntime:      "TOKEN已过期,请重新登陆",
	ErrTokenWrong:        "TOKEN不正确,请重新登陆",
	ErrTokenTypeWrong:    "TOKEN格式错误,请重新登陆",
	ErrUserNoRight:       "该用户无权限",

	ErrArtNotExists: "文章不存在",
	ErrArtExists:    "文章已经存在",

	ErrCateNameUsed:  "该分类已存在",
	ErrCateNotExists: "该分类不存在",

	ErrProfileNotExists: "主页不存在",
}

// 返回错误信息
func GetErrMsg(code int) string {
	return Err2Msg[code]
}
