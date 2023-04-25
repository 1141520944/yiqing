package mysql

import "errors"

var (
	ErrorUserExist           = errors.New("用户已存在")
	ErrorUserNotExist        = errors.New("用户不存在")
	ErrorInvalidPassword     = errors.New("用户名或密码错误")
	ErrorInvalidID           = errors.New("无效的ID")
	ErrorStudenIDExist       = errors.New("学生学号已经存在")
	ErrorStudenIDNotExist    = errors.New("学生学号不存在")
	ErrorSuperStudenExist    = errors.New("super学生已经存在")
	ErrorSuperStudenNotExist = errors.New("super学生不存在")
	ErrorStudenExist         = errors.New("open_id已经存在")
	ErrorStudenNotExist      = errors.New("open_id不存在")
	ErrorInstructorExist     = errors.New("辅导员已经存在")
	ErrorInstructorNotExist  = errors.New("辅导员不存在")
	ErrorClassExist          = errors.New("班级已经存在")
	ErrorClassNotExist       = errors.New("班级不存在")
)
