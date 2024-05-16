package main

import (
	"scripts-client/common"
	"scripts-client/controller"
)


func main() {
	common.InitConfig()
	// 连接服务器
	controller.TcpConnectionManger()
}
