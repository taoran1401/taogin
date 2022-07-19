package main

import (
	"fmt"
	"taogin/config"
	"taogin/core"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod vendor
////go:generate go mod tidy

func main() {
	//viper
	config.TAO_VIPER = core.Viper()
	//zap: 非常快的、结构化的，分日志级别的Go日志库
	config.TAO_LOG = core.Zap()
	//gorm连接数据库

	fmt.Println(config.TAO_SERVER.Zap.Director)
	//run
	core.ServerRun()
}
