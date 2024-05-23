package common

import (
	"flag"
	"strings"
)

var Conf = new(config)

// Config 定义一个结构体来存储参数
type config struct {
	Address   string
	Script    []string
	PathCheck bool
	Reconnect int
}

func init() {
	Conf = &config{}
}

func InitConfig() {
	// 解析启动参数并存储在结构体中
	flag.StringVar(&Conf.Address, "address", "127.0.0.1:7070", "Script management server address")
	flag.BoolVar(&Conf.PathCheck, "check", false, "Check whether the script file path needs to be verified")
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
