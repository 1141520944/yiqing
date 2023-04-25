package mysql

// 初始化mysql连接
import (
	"fmt"
	"webGin/models"
	"webGin/setting"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// db 数据库对象
var db *gorm.DB

// Init 初始化数据库连接
func Init(conf setting.MysqlConfig) (err error) {
	// dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DBname)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.AutoMigrate(&models.Table_class{}, &models.Table_user{},
		&models.Table_instructor{}, &models.Table_nucleicAcid{},
		&models.Table_super{}, &models.Table_admin{})
	adminCreate()
	return
}

func Close() {

}

func adminCreate() {
	var count int64
	var re *models.Table_admin
	db.Model(&models.Table_admin{}).Where("admin_id=?", 1).Count(&count).Find(re)
	if count == 0 {
		db.Model(&models.Table_admin{}).Create(&models.Table_admin{
			Username: setting.Conf.Admin.Username,
			Password: setting.Conf.Admin.Password,
			AdminID:  1,
		})
	} else {
		return
	}

}
