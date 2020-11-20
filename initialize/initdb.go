package initialize

import (
	"nav/global"
	"nav/model"
	"github.com/jinzhu/gorm"
)

// 迁移数据库
func RegisterModel(db *gorm.DB) {
	global.NLY_LOG.Info("开始迁移数据库")
	db.AutoMigrate(&model.User{})
	global.NLY_LOG.Info("迁移表 User 成功")
	global.NLY_LOG.Info("迁移数据库成功")
}


func InitDB() *gorm.DB {
	if global.NLY_CONFIG.System.DbType == "mysql" {
		// 调用mysql 初始化方法
		return InitMysql()
	}
	if global.NLY_CONFIG.System.DbType == "postgres" {
		// 调用 postgres 初始化方法<
		return InitPostgres()
	}
	if global.NLY_CONFIG.System.DbType == "sqlite" {
		// 调用 postgres 初始化方法<
		return nil
	}

	return nil
}