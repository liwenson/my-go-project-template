package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "nav/docs"
	"nav/global"
	"nav/middleware"
	"nav/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	global.NLY_LOG.Info("开始初始化路由设置")
	var Router = gin.New()
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	global.NLY_LOG.Info("使用 Zap 自定义 Logger 和 Recovery ")

	//Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	//global.NLY_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.NLY_LOG.Info("注册跨域中间件")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.NLY_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	global.NLY_LOG.Info("开始注册路由")
	ApiGroup := Router.Group("")
	router.InitUserRouers(ApiGroup) // 注册用户路由
	global.NLY_LOG.Info("全部路由注册成功")
	return Router
}
