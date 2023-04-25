package models

//gorm 创建创建表所用的结构体
import (
	"gorm.io/gorm"
)

// User 普通用户
type UserInfo struct {
	Table_class Table_class `gorm:"foreignkey:ClassID;association_foreignkey:ClassID"`
	ClassID     uint        `gorm:"column:class_id;index:idx_user_name;unique;default:'番茄炒蛋'"`
}

// Table_user 用户 -student
type Table_user struct {
	gorm.Model
	Phone         string `gorm:"comment:联系方式" json:"phone"`                                            //联系方式
	StudentName   string `gorm:"comment:学生姓名" json:"student_name"`                                     //学生姓名
	College       string `gorm:"comment:学院" json:"college"`                                            //学院
	StudentNumber string `gorm:"comment:学号;primary_key;unique" json:"student_number"`                  //学号
	OpenID        string `gorm:"comment:小程序用户的ID;index:idx_open_id;primary_key;unique" json:"open_id"` //微信小程序Open_id
	UserID        int64  `gorm:"comment:用户ID;primary_key;index:idx_user_id;unique" json:"user_id"`     //用户ID
	ClassID       int64  `gorm:"comment:班级ID" json:"class_id"`                                         //班级ID
	InstructorID  int64  `gorm:"comment:辅导ID" json:"instructor_id"`                                    //辅导ID
	TodayState    int    `gorm:"comment:今天状态" json:"today_state"`                                      //今天状态
	RoleID        int    `gorm:"comment:角色" json:"_"`                                                  //角色 -权限
	TodayTime     int64  `gorm:"comment:今天核酸时间" json:"todatTime"`                                      //今天核酸时间
}

// TableName 指定表名
func (Table_user) TableName() string {
	return "student"
}

// NucleicAcid 核酸信息
type Table_nucleicAcid struct {
	gorm.Model
	Time          int64  `gorm:"comment:核酸时间" json:"time"`                             //核酸时间
	Address       string `gorm:"comment:核酸地点" json:"address"`                          //核酸地点
	NucleicAcidID int64  `gorm:"comment:核酸ID;primary_key;unique" json:"nucle_acid_id"` //核酸ID
	UserID        int64  `gorm:"comment:核酸所有者;index:idx_user_id" json:"user_id"`       //核酸所有者
	SuperID       int64  `gorm:"comment:操作员的ID" json:"super_id"`                       //操作员的ID
	State         int    `gorm:"comment:状态" json:"state"`                              //状态
}

// TableName 指定表名
func (Table_nucleicAcid) TableName() string {
	return "nucleic_acid"
}

// Table_super 扫码用户-super学生
type Table_super struct {
	gorm.Model
	Username     string `gorm:"comment:用户名" json:"username"`                                 //用户名
	Password     string `gorm:"comment:密码" json:"password"`                                  //密码
	SuperID      int64  `gorm:"操作员ID;primary_key;index:idx_super_id;unique" json:"super_id"` //操作员ID
	InstructorID int64  `gorm:"comment:辅导员ID" json:"instructor_id"`                          //辅导员ID
}

// TableName 指定表名
func (Table_super) TableName() string {
	return "super_student"
}

// Instructor辅导员
type Table_instructor struct {
	gorm.Model
	InstructorID int64  `gorm:"comment:辅导员ID;primary_key;index:idx_instructor_id;unique" json:"instructor_id"` //辅导员ID
	Username     string `gorm:"comment:用户名" json:"username"`                                                   //用户名
	Password     string `gorm:"comment:密码" json:"password"`
	Name         string `gorm:"comment:名字" json:"name"` //密码
}

// TableName 指定表名
func (Table_instructor) TableName() string {
	return "instructor"
}

// Class 班级
type Table_class struct {
	gorm.Model
	InstructorID int64  `gorm:"comment:辅导员ID" json:"instructor_id"`                                 //辅导员ID
	ClassID      int64  `gorm:"comment:班级ID;primary_key;index:idx_class_id;unique" json:"class_id"` //班级ID
	Number       int    `gorm:"comment:人数" json:"number"`                                           //人数
	StateNumber  int    `gorm:"comment:核酸状态人数" json:"state_number"`                                 //核酸状态人数
	Grade        string `gorm:"comment:年级" json:"grade"`                                            //年级
	College      string `gorm:"comment:学院" json:"college"`                                          //学院
	Name         string `gorm:"comment:描述" json:"name"`                                             //描述
}

// TableName 指定表名
func (Table_class) TableName() string {
	return "class"
}

// admin 管理员
type Table_admin struct {
	gorm.Model
	AdminID  int64  `gorm:"comment:adminID;primary_key;unique" json:"instructor_id"` //辅导员ID
	Username string `gorm:"comment:用户名" json:"user_name"`                            //用户名
	Password string `gorm:"密码" json:"password"`                                      //密码
}

// TableName 指定表名
func (Table_admin) TableName() string {
	return "admin"
}
