package mysql

import (
	"time"
	"webGin/models"
)

// InsertOneStudentInfo 插入一条核酸记录
func InsertOneNucleicAcidInfo(p *models.Table_nucleicAcid) (err error) {
	return db.Create(p).Error
}

// GetNowTime 得到现在时间的年月日的时间戳
func GetNowTime() int64 {
	tm := time.Now().Format("2006-01-02")
	tt, _ := time.ParseInLocation("2006-01-02", tm, time.Local)
	return tt.Unix()
}

// 查找班级下的核酸信息
func SelectByClassID(p int64) (number int64, Rnumber int64, err error) {
	rm := GetNowTime()
	err = db.Model(&models.Table_user{}).Where("class_id = ? AND  today_time >= ?", p, rm).Count(&Rnumber).Error
	if err != nil {
		return 0, 0, err
	}
	err = db.Model(&models.Table_user{}).Where("class_id = ?", p).Count(&number).Error
	return
}

// 查找辅导员下的核酸信息
func SelectByInstructorId(p int64) (number int64, Rnumber int64, err error) {
	rm := GetNowTime()
	err = db.Model(&models.Table_user{}).Where("instructor_id = ? AND  today_time >= ?", p, rm).Count(&Rnumber).Error
	if err != nil {
		return 0, 0, err
	}
	err = db.Model(&models.Table_user{}).Where("instructor_id = ?", p).Count(&number).Error
	return
}

// 查找对应学生的核酸信息
func SelectByUserID(p int64) (number int64, Rnumber int64, err error) {
	rm := GetNowTime()
	err = db.Model(&models.Table_user{}).Where("user_id = ? AND  today_time >= ?", p, rm).Count(&Rnumber).Error
	if err != nil {
		return 0, 0, err
	}
	err = db.Model(&models.Table_user{}).Where("user_id = ?", p).Count(&number).Error
	return
}

// SelectoneByClassID 查找对应的Class 名称
func SelectoneByClassID(p int64) (re *models.Table_class, err error) {
	re = new(models.Table_class)
	err = db.Model(&models.Table_class{}).Where("class_id=?", p).Find(re).Error
	return re, err
}
