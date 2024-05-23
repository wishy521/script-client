package main

import (
	"fmt"
	"net"
	"scripts-client/common"
	"scripts-client/controller"
	"time"
)

func main() {
	// 初始化获取到的配置
	common.InitConfig()

	// 使用chan将TCP的内容传入到文件
	content := make(chan *string)

	go func() {
		for originalContent := range content {
			scriptInfo, scriptContent, err := controller.ExtractContent(originalContent)
			if err != nil {
				common.Log.Error("Parsing original Content failed")
				break
			}
			err = controller.WriteContentToFile(scriptInfo, scriptContent)
			if err != nil {
				common.Log.Error("write content to the file failed")
				break
			}
			common.Log.Info(fmt.Sprintf("updated file %s successfully", scriptInfo.FileInfo.Path))
		}
	}()

	for {
		conn, err := net.Dial("tcp", common.Conf.Address)
		if err != nil {
			common.Log.Error(fmt.Sprintf("connected to %s failed", common.Conf.Address))
			if common.Conf.Reconnect == 0 {
				break
			}
			time.Sleep(time.Duration(common.Conf.Reconnect) * time.Second)
			common.Log.Info(fmt.Sprintf("Attempt to reconnect %s", common.Conf.Address))
			continue
		}
		common.Log.Info(fmt.Sprintf("successfully connected to %s", conn.RemoteAddr().String()))
		controller.HandleConnection(conn, content)
	}
}
