package controller

// ResCode 定义返回值类型
type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken

	CodeStudenIDExist
	CodeStudenIDNotExist
	CodeSuperIDExist
	CodeSuperIDNotExist
	CodeClassExist
	CodeClassNotExist
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器繁忙",

	CodeStudenIDExist:    "学生学号已经被注册",
	CodeStudenIDNotExist: "学生学号没注册",
	CodeSuperIDExist:     "super学生用户名已存在",
	CodeSuperIDNotExist:  "super学生用户名不存在",
	CodeClassExist:       "班级已存在",
	CodeClassNotExist:    "班级不存在",

	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的token",
}

// GetMsg 得到对应的错误
func (r ResCode) GetMsg() string {
	msg, ok := codeMsgMap[r]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
