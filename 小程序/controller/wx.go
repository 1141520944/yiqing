package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"webGin/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thedevsaddam/gojsonq/v2"
	"go.uber.org/zap"
)

func WxGetOpenID(c *gin.Context) {
	p := new(models.WxOpen_id)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("WxGetOpenID with invalid param ", zap.Error(err))
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

	code2sessionURL := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	//调用auth.code2Session接口获取openid
	url := fmt.Sprintf(code2sessionURL, p.AppID, p.Secret, p.JsCode)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("openid is failed", zap.Error(err))
		return
	}
	json1 := gojsonq.New().FromString(string(body)).Find("openid")
	json2 := gojsonq.New().FromString(string(body)).Find("session_key")
	openId := json1.(string)
	session_key := json2.(string)
	fmt.Println("my openid is: ", openId)
	re := &models.WxResopnseOpenID{
		OpenID:     openId,
		SessionKey: session_key,
	}
	ResponseSuccess(c, re)
}
