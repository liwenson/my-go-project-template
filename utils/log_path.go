package utils

import (
	"go.uber.org/zap"
	"nav/global"
	"os"
)

// 判断文件或者文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 批量创建文件夹
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.NLY_LOG.Debug("创建日志文件夹：" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.NLY_LOG.Error("批量创建文件夹："+v, zap.Any(" error:", err))
			}
		}
	}
	return err
}
