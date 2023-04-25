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

// AddInformationSuper 添加操作员
func AddInformationSuper(c *gin.Context) {
	p := new(models.ParamSignupSuperStudent)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddInformationSuper with invalid param ", zap.Error(err))
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
	if err := logic.AddInformationSuper(p); err != nil {
		zap.L().Error("logic.AddInformationSuper() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorSuperStudenExist) {
			ResponseError(c, CodeSuperIDExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, CodeSuccess)
}

// UpdateInformationSuper 更新操作员
func UpdateInformationSuper(c *gin.Context) {
	p := new(models.ParamUpdateSuperStudent)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateInformationSuper with invalid param ", zap.Error(err))
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
	if err := logic.UpdateInformationSuper(p); err != nil {
		zap.L().Error("logic.UpdateInformationSuper() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorSuperStudenExist) {
			ResponseError(c, CodeSuperIDExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, CodeSuccess)
}

// 删除操作员 -硬删除
func DeleteInformationSuper(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//业务逻辑
	if err := logic.DeleteCompletelyInformationSuper(uid); err != nil {
		zap.L().Error("logic.DeleteCompletelyInformationSuper() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// 查询单个
func SelectOneInformationSuper(c *gin.Context) {
}

// SelectPageInformationSuper 分页查询
func SelectPageInformationSuper(c *gin.Context) {
	p := new(models.ParamPage)
	page, page_size, err := GetPageInfo(c)
	if err != nil {
		zap.L().Error("SelectPageInformationSuper() failed", zap.Error(err))
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
	total, re, err := logic.SelectPageInformationSuper(p)
	if err != nil {
		zap.L().Error("logic.SelectPageInformationSuper() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	// ResponseSuccess(c, gin.H{
	// 	"super_students": re,
	// 	"total":          total,
	// })
	ResponseSuccessLayUi(c, 0, "成功", int(total), re)
}

// LoginHandlerSuper super 学生登录
func LoginHandlerSuper(c *gin.Context) {
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
	user, token, err := logic.LoginSuper(p)
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

// SelectPageInformationSuperAll 查看所有操作员
func SelectPageInformationSuperAll(c *gin.Context) {
	p := new(models.ParamPage)
	page, page_size, err := GetPageInfo(c)
	if err != nil {
		zap.L().Error("SelectPageInformationSuper() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	p.Page = int(page)
	p.Limit = int(page_size)
	total, re, err := logic.SelectPageInformationSuperAll(p)
	if err != nil {
		zap.L().Error("logic.SelectPageInformationSuperAll() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessLayUi(c, 0, "成功", int(total), re)
}
