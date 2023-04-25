package logic

import (
	"strconv"
	"webGin/dao/mysql"
	"webGin/models"
	"webGin/pkg/jwt"
	"webGin/pkg/snowflake"
)

// AddInformationSuper 添加supper -学生记录
func AddInformationSuper(p *models.ParamSignupSuperStudent) error {
	//校验是否用户名存在
	if err := mysql.CheckSuperStudent(p.Username); err != nil {
		return err
	}
	//添加用户
	uid := snowflake.GenID()
	super := &models.Table_super{
		Username:     p.Username,
		Password:     p.Password,
		SuperID:      uid,
		InstructorID: p.InstructorID,
	}
	if err := mysql.InsertOneSuperInfo(super); err != nil {
		return err
	}
	return nil
}

// UpdateInformationSuper 更新super -学生记录
func UpdateInformationSuper(p *models.ParamUpdateSuperStudent) error {
	//校验是否用户名存在
	if err := mysql.CheckSuperStudent(p.Username); err != nil {
		return err
	}
	super := &models.Table_super{
		Username:     p.Username,
		Password:     p.Password,
		SuperID:      p.SuperID,
		InstructorID: p.InstructorID,
	}
	if err := mysql.UpdateOneSuperInfo(super); err != nil {
		return err
	}
	return nil
}

// DeleteCompletelyInformationSuper 硬删除super-学生记录
func DeleteCompletelyInformationSuper(uid int64) error {
	super := &models.Table_super{
		SuperID: uid,
	}
	if err := mysql.CDeleteOneSuperInfo(super); err != nil {
		return err
	}
	return nil
}

// SelectPageInformationSuper 分页查询super学生记录
func SelectPageInformationSuper(p *models.ParamPage) (int64, []*models.Super, error) {
	total, re, err := mysql.SelectPageSuperInfo(p.Page, p.Limit, p.Uid)
	if err != nil {
		return 0, nil, err
	}
	result := make([]*models.Super, 0, len(re))
	for _, v := range re {
		one := &models.Super{
			CreatTime:    v.Model.CreatedAt,
			UpdateTime:   v.Model.UpdatedAt,
			Username:     v.Username,
			Password:     v.Password,
			SuperID:      strconv.FormatInt(v.SuperID, 10),
			InstructorID: strconv.FormatInt(v.InstructorID, 10),
		}
		result = append(result, one)
	}
	return total, result, nil
}

// LoginSuper super用户登录
func LoginSuper(p *models.ParamLogin) (*models.Table_super, string, error) {
	t, err := mysql.LoginSuper(p)
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

// SelectPageInformationSuperAll 查询全部
func SelectPageInformationSuperAll(p *models.ParamPage) (int64, []*models.Super, error) {
	total, re, err := mysql.SelectPageSuperInfoAll(p.Page, p.Limit)
	if err != nil {
		return 0, nil, err
	}
	result := make([]*models.Super, 0, len(re))
	for _, v := range re {
		one := &models.Super{
			CreatTime:    v.Model.CreatedAt,
			UpdateTime:   v.Model.UpdatedAt,
			Username:     v.Username,
			Password:     v.Password,
			SuperID:      strconv.FormatInt(v.SuperID, 10),
			InstructorID: strconv.FormatInt(v.InstructorID, 10),
		}
		result = append(result, one)
	}
	return total, result, nil
}
