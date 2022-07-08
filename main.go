package main

import "taogin/core"

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod vendor

func main() {
	// run
	core.ServerRun()
}
