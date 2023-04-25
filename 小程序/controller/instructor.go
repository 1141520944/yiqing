package controller

import (
	"errors"
	"strconv"
	"webGin/dao/mysql"
	"webGin/logic"
	"webGin/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SingUpHandlerInstructor 辅导员注册
func SingUpHandlerInstructor(c *gin.Context) {
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SingUpHandlerInstructor with invalid param ", zap.Error(err))
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
	//业务处理
	if err := logic.SingUp(p); err != nil {
		zap.L().Error("logic.SingUp() failed", zap.Error(err))
		//依据错误的类型进行返回
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, CodeSuccess)
}

// LoginHandlerInstructor 辅导员登录
func LoginHandlerInstructor(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("LoginHandlerInstructor with invailde param", zap.Error(err))
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
	user, token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		} else {
			ResponseError(c, CodeInvalidPassword)
			return
		}
	}
	ResponseSuccess(c, gin.H{
		"userID":   strconv.FormatInt(user.InstructorID, 10), //前端js的范围:<53-1  后端int64的范围:<63-1
		"userName": user.Username,
		"token":    token,
	})
}

// UpdateInformationInstructor 更新操作员
func UpdateInformationInstructor(c *gin.Context) {
	p := new(models.ParamUpdateInstructor)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateInformationInstructor with invalid param ", zap.Error(err))
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
	if err := logic.UpdateInformationInstructor(p); err != nil {
		zap.L().Error("logic.UpdateInformationInstructor() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, CodeSuccess)
}

// DeleteInformationInstructor 删除操作员 -硬删除
func DeleteInformationInstructor(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//业务逻辑
	if err := logic.DeleteCompletelyInformationInstructor(uid); err != nil {
		zap.L().Error("logic.DeleteCompletelyInformationInstructor() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// SelectPageInformationInstructor 分页查询
func SelectPageInformationInstructor(c *gin.Context) {
	p := new(models.ParamPage)
	page, page_size, err := GetPageInfo(c)
	if err != nil {
		zap.L().Error("SelectPageInformationSuper() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	p.Page = int(page)
	p.Limit = int(page_size)
	total, re, err := logic.SelectPageInformationInstructor(p)
	if err != nil {
		zap.L().Error("logic.SelectPageInformationInstructor() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// ResponseSuccess(c, gin.H{
	// 	"instructors": re,
	// 	"total":       total,
	// })
	ResponseSuccessLayUi(c, 0, "成功", int(total), re)
}

// GetOneByInstructorIDInstructor
func GetOneByInstructorIDInstructor(c *gin.Context) {
	instructor_idStr := c.Query("instructor_id")
	instructor_id, err := strconv.ParseInt(instructor_idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	re, err := logic.GetOneByInstructorIDInstructor(instructor_id)
	if err != nil {
		zap.L().Error("logic.GetOneByInstructorIDInstructor()failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, re)
}
