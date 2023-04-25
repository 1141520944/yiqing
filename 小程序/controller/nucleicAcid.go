package controller

import (
	"webGin/logic"
	"webGin/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//核酸结果

// AddInformationNucleicAcid 添加核酸信息
func AddInformationNucleicAcid(c *gin.Context) {
	p := new(models.ParamNucleicAcidAdd)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddInformationNucleicAcid with invalid param ", zap.Error(err))
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
	p.SuperID = userID
	//业务逻辑
	if err := logic.AddInformationNucleicAcid(p); err != nil {
		zap.L().Error("logic.AddInformationNucleicAcid() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// GetClassStudentNucleicAcid
func GetClassStudentNucleicAcid(c *gin.Context) {
	p := new(models.ParamGetClassNucleicAcidNumber)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddInformationNucleicAcid with invalid param ", zap.Error(err))
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
	re, err := logic.GetClassStudentNucleicAcid(p)
	if err != nil {
		zap.L().Error("logic.GetClassStudentNucleicAcid() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, re)
}
