package controller

import (
	"errors"
	"strconv"
	"webGin/dao/mysql"
	"webGin/logic"
	"webGin/models"
	"webGin/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// GetTokenStudent 得到Token普通学生
func GetTokenStudent(c *gin.Context) {
	// userID, err := GetCyrrentUserID(c)
	// if err != nil {
	// 	zap.L().Error("GetCyrrentUserID() failed", zap.Error(err))
	// 	ResponseError(c, CodeNeedLogin)
	// 	return
	// }

}

// RegistStudent 学生注册
func RegistStudent(c *gin.Context) {
	p := new(models.ParamStudentInfoRegist)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddInformationStudent with invalid param ", zap.Error(err))
		//判断是否是校验错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, RemoveTopStruct(errs.Translate(trans)))
			return
		}
	}
	//业务逻辑
	st, err := logic.RegistInformationStudent(p)
	if err != nil {
		zap.L().Error("logic.RegistInformationStudent() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorStudenIDExist) {
			ResponseError(c, CodeStudenIDExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	//返回token和uid -生成JWT的token
	token, _ := jwt.GenToken(st.UserID, st.StudentName)
	ResponseSuccess(c, gin.H{
		"userID":   strconv.FormatInt(st.UserID, 10), //前端js的范围:<53-1  后端int64的范围:<63-1
		"userName": st.StudentName,
		"token":    token,
	})
}

// IsOrNoRegistStudent 校验学生是否登录
func IsOrNoRegistStudent(c *gin.Context) {
	p := new(models.ParamStudentInfoIsRegist)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddInformationStudent with invalid param ", zap.Error(err))
		//判断是否是校验错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, RemoveTopStruct(errs.Translate(trans)))
			return
		}
	}
	//业务逻辑
	token, student, err := logic.IsOrNoRegistStudent(p.OpenID)
	if err != nil {
		zap.L().Error("logic.IsOrNoRegistStudent() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorStudenNotExist) {
			ResponseError(c, CodeStudenIDNotExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, gin.H{
		"student": student,
		"token":   token,
	})

}

// AddInformationStudent 添加学生信息
func AddInformationStudent(c *gin.Context) {
	p := new(models.ParamStudentInfoSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddInformationStudent with invalid param ", zap.Error(err))
		//判断是否是校验错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, RemoveTopStruct(errs.Translate(trans)))
			return
		}
	}
	//依据token拿到uid
	userID, err := GetCyrrentUserID(c)
	if err != nil {
		zap.L().Error("GetCyrrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.InstructorID = userID
	//业务逻辑
	if err := logic.AddInformation(p); err != nil {
		zap.L().Error("logic.AddInformation() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorStudenIDExist) {
			ResponseError(c, CodeStudenIDExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, CodeSuccess)
}

// UpdateInformationStudent 更新学生信息
func UpdateInformationStudent(c *gin.Context) {
	p := new(models.ParamStudentInfoUpdate)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateInformationStudent with invalid param ", zap.Error(err))
		//判断是否是校验错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, RemoveTopStruct(errs.Translate(trans)))
			return
		}
	}
	//依据token拿到uid
	userID, err := GetCyrrentUserID(c)
	if err != nil {
		zap.L().Error("GetCyrrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.InstructorID = userID
	//业务逻辑
	if err := logic.UpdateInformation(p); err != nil {
		zap.L().Error("logic.UpdateInformation() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorStudenIDExist) {
			ResponseError(c, CodeStudenIDExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, CodeSuccess)
}

// DeleteInformationStudent 软删除学生信息
func DeleteInformationStudent(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//业务逻辑
	if err := logic.DeleteInformation(uid); err != nil {
		zap.L().Error("logic.DeleteInformation() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorStudenIDNotExist) {
			ResponseError(c, CodeStudenIDNotExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, CodeSuccess)
}

// DeleteCompletelyInformationStudent 完全删除学生
func DeleteCompletelyInformationStudent(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//业务逻辑
	if err := logic.DeleteCompletelyInformation(uid); err != nil {
		zap.L().Error("logic.DeleteCompletelyInformation() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorStudenIDNotExist) {
			ResponseError(c, CodeStudenIDNotExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, CodeSuccess)
}

// SelectPageInformationStudent 班级--学生
func SelectPageInformationStudent(c *gin.Context) {
	p := new(models.ParamPage)
	page, page_size, err := GetPageInfo(c)
	if err != nil {
		zap.L().Error("SelectPageInformationSuper() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	class_idStr := c.Query("class_id")
	class_id, _ := strconv.ParseInt(class_idStr, 10, 64)
	p.Limit = int(page_size)
	p.Page = int(page)
	p.Uid = class_id
	total, re, err := logic.SelectPageInformationStudent(p)
	if err != nil {
		zap.L().Error("logic.SelectPageInformationSuper() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessLayUi(c, 0, "成功", int(total), re)
}

// SelectPageInformationNucleicAcid 学生查看核酸登记记录
func SelectPageInformationNucleicAcid(c *gin.Context) {
	p := new(models.ParamPage)
	page, page_size, err := GetPageInfo(c)
	if err != nil {
		zap.L().Error("SelectPageInformationNucleicAcid() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//依据token拿到uid
	userID, err := GetCyrrentUserID(c)
	if err != nil {
		zap.L().Error("GetCyrrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.Page = int(page)
	p.Limit = int(page_size)
	p.Uid = userID
	total, re, err := logic.SelectPageInformationNucleicAcid(p)
	if err != nil {
		zap.L().Error("logic.SelectPageInformationNucleicAcid() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, gin.H{
		"nucleicAcids": re,
		"total":        total,
	})
}

// GetOneByUidStudent
func GetOneByUidStudent(c *gin.Context) {
	uidStr := c.Query("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//业务逻辑
	re, err := logic.SelectOneByUidStudent(uid)
	if err != nil {
		zap.L().Error("logic.SelectOneByUidStudent() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorStudenIDNotExist) {
			ResponseError(c, CodeStudenIDNotExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, re)
}

// GetAllClasses
func GetAllClasses(c *gin.Context) {
	uidStr := c.Query("instructor_id")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	i, c2, err2 := logic.SelectAllClassInfo(uid)
	if err2 != nil {
		ResponseError(c, CodeServerBusy)
	} else {
		ResponseSuccess(c, gin.H{
			"classes": c2,
			"total":   i,
		})
	}
}

// GetAllInstructors
func GetAllInstructors(c *gin.Context) {
	i, c2, err2 := logic.GetAllInstructors()
	if err2 != nil {
		ResponseError(c, CodeServerBusy)
	} else {
		ResponseSuccess(c, gin.H{
			"instructors": c2,
			"total":       i,
		})
	}
}
