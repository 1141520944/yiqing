package mysql

import (
	"webGin/models"
)

// CheckUserExit 查看用户是否已存在
func CheckUserExit(username string) (err error) {
	var count int64
	i := new(models.Table_instructor)
	if err = db.Where("username = ?", username).Find(&i).Count(&count).Error; err != nil {
		return err
	} else if count > 0 {
		//查询用户已存在
		return ErrorUserExist
	} else {
		return nil
	}
}

// InsertUser 插入数据
func InsertUser(i *models.Table_instructor) (err error) {
	return db.Create(i).Error
}

// Login 用户登录
func Login(p *models.ParamLogin) (user *models.Table_instructor, err error) {
	user = new(models.Table_instructor)
	var count int64
	if err = db.Where("username=?", p.Username).Find(user).Count(&count).Error; err != nil {
		return nil, err
	} else if count <= 0 {
		return nil, ErrorUserNotExist
	} else {
		return user, nil
	}
}

// UpdateOneInstructorInfo 更新一条super学生记录
func UpdateOneInstructorInfo(p *models.Table_instructor) (err error) {
	return db.Model(&models.Table_instructor{}).Where("instructor_id=?", p.InstructorID).Updates(p).Error
}

// CDeleteOneInstructorInfo 完全删除
func CDeleteOneInstructorInfo(p *models.Table_instructor) (err error) {
	return db.Unscoped().Model(&models.Table_instructor{}).Where("instructor_id=?", p.InstructorID).Delete(&models.Table_instructor{}).Error
}

// SelectPageInstructorInfo 分页查询
func SelectPageInstructorInfo(page, limit int) (int64, []*models.Table_instructor, error) {
	var total int64
	var re []*models.Table_instructor
	return total, re, db.Model(&models.Table_instructor{}).Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&re).Error
}

// GetOneByInstructorIDInstructor
func GetOneByInstructorIDInstructor(p int64) (re *models.Table_instructor, err error) {
	re = new(models.Table_instructor)
	err = db.Model(&models.Table_instructor{}).Where("instructor_id=?", p).First(re).Error
	return re, err
}

// SelectAllInstructor
func SelectAllInstructor() (int64, []*models.Table_instructor, error) {
	var total int64
	var re []*models.Table_instructor
	return total, re, db.Model(&models.Table_instructor{}).Count(&total).Find(&re).Error
}
