package models

//Page 分页查询
type ParamPage struct {
	Page  int
	Limit int
	Uid   int64
}
type ParamPagePost struct {
	Page  int   `json:"page,string" binding:"required"`
	Limit int   `json:"size,string" binding:"required"`
	Uid   int64 `json:"uid,string" binding:"required"`
}
