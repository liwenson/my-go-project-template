package core

import (
	"fmt"
	"go.uber.org/zap"
	"nav/global"
	"nav/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()
	global.NLY_LOG.Info("路由初始化成功")
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.NLY_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.NLY_LOG.Info("server run success on ", zap.String("address", address))
	global.NLY_LOG.Error(s.ListenAndServe().Error())
}
