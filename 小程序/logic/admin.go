package logic

import (
	"webGin/dao/mysql"
	"webGin/models"
	"webGin/pkg/jwt"
)

// LoginAdmin admin用户登录
func LoginAdmin(p *models.ParamLogin) (*models.Table_admin, string, error) {
	t, err := mysql.LoginAdmin(p)
	if err != nil {
		return nil, "", err
	}
	if p.Password != t.Password {
		return nil, "", mysql.ErrorInvalidPassword
	} else {
		//生成JWT的token
		token, err := jwt.GenToken(t.AdminID, t.Username)
		return t, token, err
	}
}
