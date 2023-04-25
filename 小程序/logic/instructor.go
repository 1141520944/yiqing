package logic

import (
	"strconv"
	"webGin/dao/mysql"
	"webGin/models"
	"webGin/pkg/jwt"
	"webGin/pkg/snowflake"
)

// SingUp 用户注册业务逻辑
func SingUp(p *models.ParamSignUp) error {
	//判断用户存不存在
	err := mysql.CheckUserExit(p.Username)
	if err != nil {
		//数据库查询错误
		return err
	}
	//生成UID
	id := snowflake.GenID()
	//构造一个User的实例
	user := models.Table_instructor{
		InstructorID: id,
		Username:     p.Username,
		Password:     p.Password,
	}
	//密码加密
	//保存进数据库
	return mysql.InsertUser(&user)
}

// Login 用户登录
func Login(p *models.ParamLogin) (*models.Table_instructor, string, error) {
	t, err := mysql.Login(p)
	if err != nil {
		return nil, "", err
	}
	if p.Password != t.Password {
		return nil, "", mysql.ErrorInvalidPassword
	} else {
		//生成JWT的token
		token, err := jwt.GenToken(t.InstructorID, t.Username)
		return t, token, err
	}
}

// UpdateInformationInstructor 更新辅导员数据
func UpdateInformationInstructor(p *models.ParamUpdateInstructor) error {
	//校验是否用户名存在
	if err := mysql.CheckUserExit(p.Username); err != nil {
		return err
	}
	instructor := &models.Table_instructor{
		Name:         p.Name,
		InstructorID: p.InstructorID,
		Username:     p.Username,
		Password:     p.Password,
	}
	if err := mysql.UpdateOneInstructorInfo(instructor); err != nil {
		return err
	}
	return nil
}

// DeleteCompletelyInformationInstructor 硬删除super-学生记录
func DeleteCompletelyInformationInstructor(uid int64) error {
	instructor := &models.Table_instructor{
		InstructorID: uid,
	}
	if err := mysql.CDeleteOneInstructorInfo(instructor); err != nil {
		return err
	}
	return nil
}

// SelectPageInformationInstructor 分页查询Instructor记录
func SelectPageInformationInstructor(p *models.ParamPage) (int64, []*models.Instructor, error) {
	total, re, err := mysql.SelectPageInstructorInfo(p.Page, p.Limit)
	if err != nil {
		return 0, nil, err
	}
	result := make([]*models.Instructor, 0, len(re))
	for _, v := range re {
		one := &models.Instructor{
			CreatTime:    v.Model.CreatedAt,
			UpdateTime:   v.Model.UpdatedAt,
			InstructorID: strconv.FormatInt(v.InstructorID, 10),
			Username:     v.Username,
			Password:     v.Password,
			Name:         v.Name,
		}
		result = append(result, one)
	}
	return total, result, nil
}

// GetOneByInstructorIDInstructor
func GetOneByInstructorIDInstructor(p int64) (*models.Instructor, error) {
	re, err := mysql.GetOneByInstructorIDInstructor(p)
	if err != nil {
		return nil, err
	} else {
		result := &models.Instructor{
			CreatTime:    re.CreatedAt,
			UpdateTime:   re.UpdatedAt,
			InstructorID: strconv.FormatInt(re.InstructorID, 10),
			Username:     re.Username,
			Password:     re.Password,
			Name:         re.Name,
		}
		return result, nil
	}
}
