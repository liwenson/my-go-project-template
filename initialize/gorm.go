package initialize

import (
	"nav/global"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/zap"
	//"note/utils"
)




// 构造 数据库初始化函数
func InitMysql() *gorm.DB {
	global.NLY_LOG.Info("开始初始化 Mysql 数据库")
	hostpath := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		global.NLY_CONFIG.Mysql.Username,
		global.NLY_CONFIG.Mysql.Password,
		global.NLY_CONFIG.Mysql.Path,
		global.NLY_CONFIG.Mysql.Dbname,
		global.NLY_CONFIG.Mysql.Config,
	)
	db, err := gorm.Open("mysql", hostpath)
	if err != nil {
		global.NLY_LOG.Info("数据库连接失败", zap.String("hostpath", hostpath))
		return nil
	} else {
		global.NLY_LOG.Info("mysql 连接成功: ", zap.String("hostpath", hostpath))
		db.DB().SetMaxIdleConns(global.NLY_CONFIG.Mysql.MaxIdleConns)  //最大打开的连接数
		db.DB().SetMaxOpenConns(global.NLY_CONFIG.Mysql.MaxOpenConns)  //设置最大闲置个数
		db.LogMode(true)  // 开启debug 日志
		return db
	}
}


// 构造 数据库初始化函数
func InitPostgres() *gorm.DB {
	global.NLY_LOG.Info("开始初始化数据库")
	hostname := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		global.NLY_CONFIG.Postgres.Host,
		global.NLY_CONFIG.Postgres.Port,
		global.NLY_CONFIG.Postgres.Username,
		global.NLY_CONFIG.Postgres.DataBase,
		global.NLY_CONFIG.Postgres.Password,
		global.NLY_CONFIG.Postgres.Sslmode,
	)
	db, err := gorm.Open("postgres", hostname)
	if err != nil {
		global.NLY_LOG.Info("数据库连接失败", zap.String("hostname", hostname))
		return nil
	} else {
		global.NLY_LOG.Info("postgres启动成功: ", zap.String("hostname", hostname))
		db.DB().SetMaxIdleConns(global.NLY_CONFIG.Postgres.MaxIdleConns)
		db.DB().SetMaxOpenConns(global.NLY_CONFIG.Postgres.MaxOpenConns)
		return db
	}
}
