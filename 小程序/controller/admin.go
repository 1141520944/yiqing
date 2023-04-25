package controller

import (
	"errors"
	"webGin/dao/mysql"
	"webGin/logic"
	"webGin/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// LoginHandlerInstructor admin登录
func LoginHandlerAdmin(c *gin.Context) {
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
	user, token, err := logic.LoginAdmin(p)
	if err != nil {
		zap.L().Error("logic.LoginAdmin() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		} else {
			ResponseError(c, CodeInvalidPassword)
			return
		}
	}
	ResponseSuccess(c, gin.H{
		"userID":   user.AdminID, //前端js的范围:<53-1  后端int64的范围:<63-1
		"userName": user.Username,
		"token":    token,
	})
}
