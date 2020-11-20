package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"nav/global"
	"nav/model"
	//"nav/utils"
	"os"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.NLY_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "postgres":
		return GormPostgreSQL()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

// MysqlTables 注册数据库表专用
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		global.NLY_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.NLY_LOG.Info("注册数据表成功 ")
	global.NLY_LOG.Info("register table success")
}

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.NLY_CONFIG.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	global.NLY_LOG.Info("连接"+ global.NLY_CONFIG.System.DbType + " 数据库")
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.NLY_LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}


// GormMysql 初始化Mysql数据库
func GormPostgreSQL() *gorm.DB {
	m := global.NLY_CONFIG.Postgres
	dsn := "host="+m.Host+" port="+m.Port+" user="+ m.Username+" password="+m.Password+" dbname="+m.DataBase+" sslmode="+m.Sslmode+" TimeZone=Asia/Shanghai"
	global.NLY_LOG.Info("连接"+ global.NLY_CONFIG.System.DbType + " 数据库")
	if db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{}); err != nil{
		global.NLY_LOG.Error("postgres启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	}else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}

}


// GormMysql 初始化Mysql数据库
func GormSqlite() *gorm.DB {
	m := global.NLY_CONFIG.Sqlite
	global.NLY_LOG.Info("连接 "+ global.NLY_CONFIG.System.DbType + " 数据库")
	if db, err := gorm.Open(sqlite.Open(m.Path), &gorm.Config{}); err != nil{
		global.NLY_LOG.Error("Sqlite启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	}else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}

}



// gormConfig 根据配置决定是否开启日志
func gormConfig(mod bool) *gorm.Config {
	if mod {
		return &gorm.Config{
			Logger:  logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		return &gorm.Config{
			Logger:  logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
}
