package controller

//接口文档用到的model

//_ResponsePostList
type _ResponseAny struct {
	Code    ResCode       `json:"code"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}
