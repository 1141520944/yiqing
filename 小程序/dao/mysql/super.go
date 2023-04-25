package mysql

import "webGin/models"

//CheckSuperStudent 检验用户名是否存在
func CheckSuperStudent(username string) (err error) {
	var Count int64
	err = db.Model(&models.Table_super{}).Where("username =?", username).Count(&Count).Error
	if err != nil {
		return err
	} else if Count > 0 {
		return ErrorSuperStudenExist
	} else {
		return
	}
}

//InsertOneSuperInfo 插入一条super学生记录
func InsertOneSuperInfo(p *models.Table_super) (err error) {
	return db.Create(p).Error
}

//UpdateOneSuperInfo 更新一条super学生记录
func UpdateOneSuperInfo(p *models.Table_super) (err error) {
	return db.Model(&models.Table_super{}).Where("super_id=?", p.SuperID).Updates(p).Error
}

//CDeleteOneSuperInfo 完全删除
func CDeleteOneSuperInfo(p *models.Table_super) (err error) {
	return db.Unscoped().Where("super_id=?", p.SuperID).Delete(&models.Table_super{}).Error
}

// DBConn.Model(&model.WorkFlow{}).Where("name like ?", "%"+name+"%").Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&workFlow).Error
//SelectPageSuperInfo 分页查询
func SelectPageSuperInfo(page, limit int, uid int64) (int64, []*models.Table_super, error) {
	var total int64
	var re []*models.Table_super
	return total, re, db.Model(&models.Table_super{}).Where("instructor_id=?", uid).Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&re).Error
}

//LoginSuper super学生登录
func LoginSuper(p *models.ParamLogin) (user *models.Table_super, err error) {
	user = new(models.Table_super)
	var count int64
	if err = db.Where("username=?", p.Username).Find(user).Count(&count).Error; err != nil {
		return nil, err
	} else if count <= 0 {
		return nil, ErrorUserNotExist
	} else {
		return user, nil
	}
}

//SelectPageSuperInfoAll
func SelectPageSuperInfoAll(page, limit int) (int64, []*models.Table_super, error) {
	var total int64
	var re []*models.Table_super
	return total, re, db.Model(&models.Table_super{}).Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&re).Error
}
