package models

//参数

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}

// ParamLogin 登录参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//ParamUpdateInstructor 更新Instructor学生
type ParamUpdateInstructor struct {
	Name         string `json:"name" binding:"required"`
	Username     string `json:"username" binding:"required"`             //用户名
	Password     string `json:"password" binding:"required"`             //密码
	InstructorID int64  `json:"instructor_id,string" binding:"required"` //辅导员ID
}

// ParamStudentInfoSignUp 添加学生信息
type ParamStudentInfoSignUp struct {
	Phone         string `json:"phone" binding:"required"`           //联系方式
	StudentName   string `json:"student_name" binding:"required"`    //学生姓名
	College       string `json:"college" binding:"required"`         //学院
	StudentNumber string `json:"stdudent_number" binding:"required"` //学号
	ClassID       int64  `json:"class_id,string" binding:"required"` //班级ID
	InstructorID  int64  `json:"instructor_id,string"`               //辅导ID
}

// ParamStudentInfoRegist 注册学生信息
type ParamStudentInfoRegist struct {
	OpenID        string `json:"open_id" binding:"required"`              //微信open_id
	Phone         string `json:"phone" binding:"required"`                //联系方式
	StudentName   string `json:"student_name" binding:"required"`         //学生姓名
	College       string `json:"college" binding:"required"`              //学院
	StudentNumber string `json:"stdudent_number" binding:"required"`      //学号
	ClassID       int64  `json:"class_id,string" binding:"required"`      //班级ID
	InstructorID  int64  `json:"instructor_id,string" binding:"required"` //辅导ID
}

// ParamStudentInfoIsRegist 注册学生信息
type ParamStudentInfoIsRegist struct {
	OpenID string `json:"open_id" binding:"required"` //微信open_id
}

// ParamStudentInfoUpdate 更新学生信息
type ParamStudentInfoUpdate struct {
	// 157791956242432
	Phone         string `json:"phone" binding:"required"`           //联系方式
	StudentName   string `json:"student_name" binding:"required"`    //学生姓名
	College       string `json:"college" binding:"required"`         //学院
	StudentNumber string `json:"stdudent_number" binding:"required"` //学号
	UserID        int64  `json:"u_id,string" binding:"required"`
	ClassID       int64  `json:"class_id,string" binding:"required"` //班级ID
	InstructorID  int64  `json:"instructor_id,string"`               //辅导ID
}

//ParamNucleicAcidAdd 添加核酸记录
type ParamNucleicAcidAdd struct {
	Address string `json:"adderss" binding:"required"`        //核酸地点
	UserID  int64  `json:"user_id,string" binding:"required"` //核酸所有者
	SuperID int64  `json:"super_id"`                          //操作员ID
	State   int    `json:"state"`                             //状态
}

//ParamSuperStudent 添加super学生
type ParamSignupSuperStudent struct {
	Username     string `json:"username" binding:"required"` //用户名
	Password     string `json:"password" binding:"required"` //密码
	InstructorID int64  `json:"instructor_id,string"`        //辅导员ID
}

//ParamUpdateSuperStudent 更新super学生
type ParamUpdateSuperStudent struct {
	Username     string `json:"username" binding:"required"`        //用户名
	Password     string `json:"password" binding:"required"`        //密码
	SuperID      int64  `json:"super_id,string" binding:"required"` //操作员ID
	InstructorID int64  `json:"instructor_id,string"`               //辅导员ID
}

//ParamAddClass 增加班级
type ParamAddClass struct {
	InstructorID int64  `json:"instructor_id"`              //辅导员ID
	Grade        string `json:"grade" binding:"required"`   //年级
	College      string `json:"college" binding:"required"` //学院
	Name         string `json:"name" binding:"required"`    //描述
}

//ParamUpdateClass 更新班级
type ParamUpdateClass struct {
	InstructorID int64  `json:"instructor_id" `                     //辅导员ID
	ClassID      int64  `json:"class_id,string" binding:"required"` //班级ID
	Grade        string `json:"grade" binding:"required"`           //年级
	College      string `json:"college" binding:"required"`         //学院
	Name         string `json:"name" binding:"required"`            //描述
}

//utf8mb4
//utf8mb4_0900_ai_ci
// ParamGetClassNucleicAcidNumber 查看对应班级学生的核酸人数
type ParamGetClassNucleicAcidNumber struct {
	StudentName   string `json:"student_name" binding:"required"`    //学生姓名
	College       string `json:"college" binding:"required"`         //学院
	StudentNumber string `json:"stdudent_number" binding:"required"` //学号
	UserID        int64  `json:"u_id,string" binding:"required"`     //userID
	ClassID       int64  `json:"class_id,string" binding:"required"` //班级ID
	InstructorID  int64  `json:"instructor_id,string"`               //辅导ID
}

//WxOpen_id 获得用户的open_id
type WxOpen_id struct {
	AppID  string `json:"app_id" binding:"required"`
	Secret string `json:"secret" binding:"required"`
	Url    string `json:"url" binding:"required"`
	JsCode string `json:"code" binding:"required"`
	Method string `json:"method" binding:"required"`
}
