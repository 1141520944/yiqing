package mysql

import "webGin/models"

// LoginAdmin 用户登录
func LoginAdmin(p *models.ParamLogin) (user *models.Table_admin, err error) {
	user = new(models.Table_admin)
	var count int64
	if err = db.Model(&models.Table_admin{}).Where("username=?", p.Username).Find(user).Count(&count).Error; err != nil {
		return nil, err
	} else if count <= 0 {
		return nil, ErrorUserNotExist
	} else {
		return user, nil
	}
}
