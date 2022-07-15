package core

import (
	"go.uber.org/zap"
	"os"
	"taogin/config"
	"taogin/core/utils"
)

func Zap() *zap.Logger {
	//检查目录
	ok, _ := utils.DirExists(config.TAO_SERVER.Zap.Director)
	if !ok {
		//创建
		os.Mkdir(config.TAO_SERVER.Zap.Director, 0777)
	}
	logger, _ := zap.NewProduction()
	return logger
}
