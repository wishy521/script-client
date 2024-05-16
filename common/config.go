package common

import (
	"flag"
	"strings"
)

var Conf = new(config)

// Config 定义一个结构体来存储参数
type config struct {
	Address    string
	Script     []string
	Reconnect  int
}

func init() {
	Conf = &config{} // 为结构体实例分配内存
}

func InitConfig() {
	// 解析启动参数并存储在结构体中
	flag.StringVar(&Conf.Address, "address", "127.0.0.1:7070", "Script management server address")
	flag.IntVar(&Conf.Reconnect, "reconnect", 10, "After disconnecting from the server, retry connection interval")

	// 解析路径参数
	scriptString := flag.String("script", "shell.sh", "Script file path")
	if strings.Contains(*scriptString, ",") {
		script := strings.Split(*scriptString, ",")
		Conf.Script = script
	} else {
		script := []string{*scriptString}
		Conf.Script = script
	}
	flag.Parse()
}
