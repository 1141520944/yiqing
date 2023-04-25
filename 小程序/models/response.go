package models

import (
	"time"
)

type Student struct {
	CreatTime  time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	// DeleteTime    time.Time `json:"delete_time"`
	Phone         string    `json:"phone"`          //联系方式
	StudentName   string    `json:"student_name"`   //学生姓名
	College       string    `json:"college"`        //学院
	StudentNumber string    `json:"student_number"` //学号
	OpenID        string    `json:"open_id"`        //微信小程序Open_id
	UserID        string    `json:"user_id"`        //用户ID
	ClassID       string    `json:"class_id"`       //班级ID
	InstructorID  string    `json:"instructor_id"`  //辅导ID
	TodayState    int       `json:"today_state"`    //今天状态
	RoleID        int       `json:"role_id"`        //角色 -权限
	TodayTime     time.Time `json:"todatTime"`
}

// super 学生
type SuperStudent struct {
	Username     string `json:"username"`      //用户名
	Password     string `json:"password"`      //密码
	SuperID      string `json:"super_id"`      //操作员ID
	InstructorID string `json:"instructor_id"` //辅导员ID
}

// 核酸的记录信息
type NucleicAcidNumberInfo struct {
	ClassName         string `json:"class_name"`                //班级名称
	ClassNumber       int64  `json:"class_number,string"`       //班级总人数
	ClassRNumber      int64  `json:"class_rnumber,string"`      //今天已经核酸人数
	InstructorNumber  int64  `json:"instructor_number,string"`  //辅导员下总人数
	InstructorRNumber int64  `json:"instructor_rnumber,string"` //辅导员下已核酸人数
	StudentName       string `json:"student_name"`              //学生姓名
	StudentIDNumber   string `json:"student_id_number"`         //学生学号
	StudentCollege    string `json:"student_college"`           //学院
	StudentGard       string `json:"student_gard"`              //年级
	StudentNumber     int64  `json:"student_number,string"`     //学生是否存在
	StudentRNumber    int64  `json:"student_rnumber,string"`    //该学生是否核酸
}

type Super struct {
	CreatTime    time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
	Username     string    `json:"username"` //用户名
	Password     string    `json:"password"` //密码
	SuperID      string    `json:"super_id"` //操作员ID
	InstructorID string    `json:"instructor_id"`
}

type Class struct {
	CreatTime    time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
	InstructorID string    `json:"instructor_id"` //辅导员ID
	ClassID      string    `json:"class_id"`      //班级ID
	Number       int       `json:"number"`        //人数
	StateNumber  int       `json:"state_number"`  //核酸状态人数
	Grade        string    `json:"grade"`         //年级
	College      string    `json:"college"`       //学院
	Name         string    `json:"name"`
}

type NucleicAcid struct {
	CreatTime     time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
	Time          string    `json:"time"`          //核酸时间
	Address       string    `json:"address"`       //核酸地点
	NucleicAcidID string    `json:"nucle_acid_id"` //核酸ID
	UserID        string    `json:"user_id"`       //核酸所有者
	SuperID       string    `json:"super_id"`      //操作员的ID
	State         int       `json:"state"`
}

type Instructor struct {
	CreatTime    time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
	InstructorID string    `json:"instructor_id"` //辅导员ID
	Username     string    `json:"username"`      //用户名
	Password     string    `json:"password"`
	Name         string    `json:"name"`
}

type WxResopnseOpenID struct {
	OpenID     string `json:"open_id"`
	SessionKey string `json:"session_key"`
}

type LayuiData struct {
	// "code": 0,
	// "msg": "",
	// "count": 1000,
	// "data": [
	// Code
}
