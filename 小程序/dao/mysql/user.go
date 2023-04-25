package mysql

import (
	"fmt"
	"webGin/models"
)

// CheckStudentExistState 查看学生的学号是否已经存在
func CheckStudentExistState(p string) (err error) {
	student := new(models.Table_user)
	var Count int64
	if err = db.Model(&models.Table_user{}).Where("student_number=?", p).Find(student).Count(&Count).Error; err != nil {
		return err
	} else if Count > 0 {
		return ErrorStudenIDExist
	} else {
		return
	}
}

// InsertOneStudentInfo 插入一条学生记录
func InsertOneStudentInfo(p *models.Table_user) (err error) {
	return db.Model(&models.Table_user{}).Omit("today_time").Create(p).Error
}

// UpdateOneStudentInfo 更新一条学生记录
func UpdateOneStudentInfo(p *models.Table_user) (err error) {
	return db.Model(&models.Table_user{}).Where("user_id", p.UserID).Omit("today_time").Updates(p).Error
}

// SelectByUid 依据uid查找 -软
func SelectByUid(uid int64) (err error) {
	var Count int64
	if err = db.Model(&models.Table_user{}).Where("user_id", uid).Count(&Count).Error; err != nil {
		return err
	} else if Count <= 0 {
		return ErrorStudenIDNotExist
	} else {
		return
	}
}

// DeleteByUid 依据uid删除 软删除
func DeleteByUid(p *models.Table_user) (err error) {
	return db.Model(&models.Table_user{}).Where("user_id=?", p.UserID).Delete(&models.Table_user{}).Error
}

// DeleteCompleteByUid  依据uid删除 完全删除
func DeleteCompleteByUid(p *models.Table_user) (err error) {
	return db.Unscoped().Model(&models.Table_user{}).Where("user_id=?", p.UserID).Delete(&models.Table_user{}).Error
}

// RegistOneStudentInfo
func RegistOneStudentInfo(p *models.Table_user) (err error) {
	return db.Model(&models.Table_user{}).Omit("today_time").Create(p).Error
}

// SelectByStudentOpenid 查找是否注册过
func SelectByStudentOpenid(p *models.Table_user) (*models.Table_user, error) {
	var count int64
	err := db.Model(&models.Table_user{}).Where("open_id=?", p.OpenID).Count(&count).Find(p).Error
	if err != nil {
		return nil, err
	} else if count <= 0 {
		return nil, ErrorStudenNotExist
	} else {
		return p, nil
	}
}

// SelectPageStudentInfo
func SelectPageStudentInfo(page, limit int, uid int64) (int64, []*models.Table_nucleicAcid, error) {
	var total int64
	var re []*models.Table_nucleicAcid
	return total, re, db.Model(&models.Table_nucleicAcid{}).Where("user_id=?", uid).Count(&total).Limit(limit).Offset((page - 1) * limit).Order("id desc").Find(&re).Error
}

// UpdateNucleAcideState 修改核酸相关的状态
func UpdateNucleAcideState(p *models.Table_nucleicAcid) (err error) {
	user := &models.Table_user{
		TodayState: 1,
		TodayTime:  p.Time,
		UserID:     p.UserID,
	}
	return db.Model(&models.Table_user{}).Where("user_id", p.UserID).Updates(user).Error
}

// SelectPageStudentClassInfo
func SelectPageStudentClassInfo(page, limit int, uid int64) (int64, []*models.Table_user, error) {
	var total int64
	var re []*models.Table_user
	return total, re, db.Model(&models.Table_user{}).Where("class_id=?", uid).Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&re).Error
}

// SelectByOpenIDStudent 依据openid查找
func SelectByOpenIDStudent(p string) (err error) {
	var count int64
	db.Model(&models.Table_user{}).Where("open_id=?", p).Count(&count)
	fmt.Printf("count: %v\n", count)
	if count > 0 {
		return ErrorStudenExist
	} else {
		return
	}
}

// UpdateByOpenIDStudent 更新字段
func UpdateByOpenIDStudent(p *models.Table_user) (re *models.Table_user, err error) {
	re = new(models.Table_user)
	return re, db.Model(&models.Table_user{}).Where("open_id=?", p.OpenID).Updates(p).Find(re).Error
}

// SelectOneByUidStudent
func SelectOneByUidStudent(uid int64) (re *models.Table_user, err error) {
	var Count int64
	re = new(models.Table_user)
	if err = db.Model(&models.Table_user{}).Where("user_id", uid).Count(&Count).First(re).Error; err != nil {
		return nil, err
	} else if Count <= 0 {
		return nil, ErrorStudenIDNotExist
	} else {
		return re, nil
	}
}
