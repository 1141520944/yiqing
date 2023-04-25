package mysql

import "webGin/models"

//CheckClass 检查班级是否重复
func CheckClass(p string) (err error) {
	var Count int64
	err = db.Model(&models.Table_class{}).Where("name = ?", p).Count(&Count).Error
	if err != nil {
		return err
	} else if Count > 0 {
		return ErrorClassExist
	} else {
		return
	}
}

//InsertOneClassInfo 创建一条class记录
func InsertOneClassInfo(p *models.Table_class) (err error) {
	return db.Create(p).Error
}

//UpdateOneClassInfo 更新一条class记录
func UpdateOneClassInfo(p *models.Table_class) (err error) {
	return db.Model(&models.Table_class{}).Where("class_id=?", p.ClassID).Updates(p).Error
}

//GetOneByClassIDClass
func GetOneByClassIDClass(p int64) (re *models.Table_class, err error) {
	re = new(models.Table_class)
	err = db.Model(&models.Table_class{}).Where("class_id=?", p).First(re).Error
	return re, err
}

//CDeleteOneStudentInfo 完全删除
func CDeleteOneStudentInfo(p *models.Table_class) (err error) {
	return db.Unscoped().Model(&models.Table_class{}).Where("class_id=?", p.ClassID).Delete(&models.Table_class{}).Error
}

// DBConn.Model(&model.WorkFlow{}).Where("name like ?", "%"+name+"%").Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&workFlow).Error
//SelectPageClassInfo 分页查询
func SelectPageClassInfo(page, limit int, uid int64) (int64, []*models.Table_class, error) {
	var total int64
	var re []*models.Table_class
	return total, re, db.Model(&models.Table_class{}).Where("instructor_id=?", uid).Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&re).Error
}

//SelectAllClassInfo
func SelectAllClassInfo(uid int64) (int64, []*models.Table_class, error) {
	var total int64
	var re []*models.Table_class
	return total, re, db.Model(&models.Table_class{}).Where("instructor_id=?", uid).Count(&total).Find(&re).Error
}
