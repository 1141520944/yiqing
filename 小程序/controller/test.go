package controller

import "github.com/gin-gonic/gin"

//@Description 用户登录
//@Summary 获取账号进行登录
//@Accept multipart/form-data
//@Produce application/json
// @Param Authorization header string flase "Bearer 用户令牌"
// @Param object query models.ParamLogin flase "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseAny
// @Router /api/v1/hello [get]
func Hello(c *gin.Context) {
	ResponseSuccess(c, "成功")
}
