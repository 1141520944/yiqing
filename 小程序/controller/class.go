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

// AddInformationClass 添加班级
func AddInformationClass(c *gin.Context) {
	p := new(models.ParamAddClass)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddInformationClass with invalid param ", zap.Error(err))
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
	if err := logic.AddInformationClass(p); err != nil {
		zap.L().Error("logic.AddInformationClass() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorClassExist) {
			ResponseError(c, CodeClassExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, CodeSuccess)
}

// UpdateInformationClass 更新班级
func UpdateInformationClass(c *gin.Context) {
	p := new(models.ParamUpdateClass)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddInformationClass with invalid param ", zap.Error(err))
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
	if err := logic.UpdateInformationClass(p); err != nil {
		zap.L().Error("logic.AddInformationClass() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorClassExist) {
			ResponseError(c, CodeClassExist)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, CodeSuccess)
}

// DeleteInformationClass 删除class -硬删除
func DeleteInformationClass(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//业务逻辑
	if err := logic.DeleteCompletelyInformationClass(uid); err != nil {
		zap.L().Error("logic.DeleteCompletelyInformationClass() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// SelectPageInformationClass 分页查询
func SelectPageInformationClass(c *gin.Context) {
	p := new(models.ParamPage)
	page, page_size, err := GetPageInfo(c)
	if err != nil {
		zap.L().Error("SelectPageInformationClass() failed", zap.Error(err))
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
	total, re, err := logic.SelectPageInformationClass(p)
	if err != nil {
		zap.L().Error("logic.SelectPageInformationClass() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	// ResponseSuccess(c, gin.H{
	// 	"classes": re,
	// 	"total":   total,
	// })
	ResponseSuccessLayUi(c, 0, "成功", int(total), re)
}

// GetOneByClassIDClass
func GetOneByClassIDClass(c *gin.Context) {
	class_idStr := c.Query("class_id")
	class_id, err := strconv.ParseInt(class_idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	re, err := logic.GetOneByClassIDClass(class_id)
	if err != nil {
		zap.L().Error("logic.GetOneByClassIDClass()failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, re)
}
