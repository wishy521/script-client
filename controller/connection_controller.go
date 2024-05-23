package controller

import (
	"net"
	"scripts-client/common"
)

// HandleConnection TCP连接管理
func HandleConnection(conn net.Conn, content chan<- *string) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			break
		}
		if string(buffer[:bytesRead]) == "START" {
			// 更新文件
			common.Log.Info("Received start updating identifier")
			buffer := make([]byte, 1024)
			for {
				bytesRead, err := conn.Read(buffer)
				if err != nil {
					common.Log.Error("Error reading from connection:")
					break
				}
				// 检查是否收到结束标识符
				if string(buffer[:bytesRead]) == "END" {
					common.Log.Info("Received end update identifier")
					break
				}
				contentString := string(buffer[:bytesRead])
				content <- &contentString
			}
			common.Log.Info("Received content successfully")
		}
	}
}
