package logic

import (
	"strconv"
	"webGin/dao/mysql"
	"webGin/models"
	"webGin/pkg/snowflake"
)

// AddInformationClass 添加班级
func AddInformationClass(p *models.ParamAddClass) error {
	//校验是否用户名存在
	if err := mysql.CheckClass(p.Name); err != nil {
		return err
	}
	//添加用户
	uid := snowflake.GenID()
	class := &models.Table_class{
		ClassID:      uid,
		Name:         p.Name,
		College:      p.College,
		Grade:        p.Grade,
		InstructorID: p.InstructorID,
	}
	if err := mysql.InsertOneClassInfo(class); err != nil {
		return err
	}
	return nil
}

// UpdateInformationClass 更新班级
func UpdateInformationClass(p *models.ParamUpdateClass) error {
	//校验是否用户名存在
	if err := mysql.CheckClass(p.Name); err != nil {
		return err
	}
	//添加用户
	class := &models.Table_class{
		ClassID:      p.ClassID,
		Name:         p.Name,
		College:      p.College,
		Grade:        p.Grade,
		InstructorID: p.InstructorID,
	}
	if err := mysql.UpdateOneClassInfo(class); err != nil {
		return err
	}
	return nil
}

// DeleteCompletelyInformationClass 硬删除class记录
func DeleteCompletelyInformationClass(uid int64) error {
	class := &models.Table_class{
		ClassID: uid,
	}
	if err := mysql.CDeleteOneStudentInfo(class); err != nil {
		return err
	}
	return nil
}

// SelectPageInformationClass 分页查询class记录
func SelectPageInformationClass(p *models.ParamPage) (int64, []*models.Class, error) {
	total, re, err := mysql.SelectPageClassInfo(p.Page, p.Limit, p.Uid)
	if err != nil {
		return 0, nil, err
	}
	result := make([]*models.Class, 0, len(re))
	for _, v := range re {
		one := &models.Class{
			CreatTime:    v.Model.CreatedAt,
			UpdateTime:   v.Model.UpdatedAt,
			InstructorID: strconv.FormatInt(v.InstructorID, 10),
			ClassID:      strconv.FormatInt(v.ClassID, 10),
			Number:       v.Number,
			StateNumber:  v.StateNumber,
			Grade:        v.Grade,
			College:      v.College,
			Name:         v.Name,
		}
		result = append(result, one)
	}
	return total, result, nil
}

// GetOneByClassIDClass
func GetOneByClassIDClass(p int64) (*models.Class, error) {
	re, err := mysql.GetOneByClassIDClass(p)
	if err != nil {
		return nil, err
	} else {
		result := &models.Class{
			CreatTime:    re.CreatedAt,
			UpdateTime:   re.UpdatedAt,
			InstructorID: strconv.FormatInt(re.InstructorID, 10),
			ClassID:      strconv.FormatInt(re.ClassID, 10),
			Number:       re.Number,
			StateNumber:  re.StateNumber,
			Grade:        re.Grade,
			College:      re.College,
			Name:         re.Name,
		}
		return result, nil
	}
}
