package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nav/config"
)

// 配置全局变量
var (
	NLY_DB      *gorm.DB
	NLY_LOG     *zap.Logger
	NLY_CONFIG  config.Server
	NLY_NIL_RES interface{}
)
