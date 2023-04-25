package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	{
		"code":1001,//程序中的错误码
		"msg":xx,提示信息
		"data":{},//数据
	}
*/
type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// ResponseError 返回错误类型
func ResponseError(c *gin.Context, code ResCode) {
	resp := &Response{
		Code: code,
		Msg:  code.GetMsg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, resp)
}

// ResponseSuccess 返回请求成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.GetMsg(),
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}

// 自定义的返回错误
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	resp := &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, resp)
}

// ResponseSuccessLayUi 返回Layui 数据
func ResponseSuccessLayUi(c *gin.Context, code int, msg string, count int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"count": count,
		"data":  data,
		"msg":   msg,
	})
}
