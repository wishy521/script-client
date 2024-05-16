package controller

import (
	"fmt"
	"net"
	"os"
	"scripts-client/common"
	"time"
)

func updateFile(conn net.Conn) {
	// 删除已存在的文件（如果有），然后创建新文件
	err := os.Remove(common.Conf.Script[0])
	if err != nil && !os.IsNotExist(err) {
		common.Log.Error("Error removing existing file:", err.Error())
		return
	}

	file, err := os.Create(common.Conf.Script[0])
	if err != nil {
		common.Log.Error("Error creating file:", err.Error())
		return
	}
	defer file.Close()

	// 从连接中读取数据并写入文件，直到收到结束标识符
	buffer := make([]byte, 1024)
	for {
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			common.Log.Error("Error reading from connection:", err.Error())
			break
		}
		// 检查是否收到结束标识符
		if string(buffer[:bytesRead]) == "END" {
			common.Log.Info("Received end update identifier")
			break
		}
		// 写入文件
		_, err = file.Write(buffer[:bytesRead])
		//fmt.Println(string(buffer[:bytesRead]))
		if err != nil {
			common.Log.Error("Error writing to file:", err.Error())
			break
		}
	}
	common.Log.Info("File received successfully")
}


func TcpConnectionManger() {
	for  {
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

		// 接收服务器的消息
		buffer := make([]byte, 1024)
		for {
			bytesRead, err := conn.Read(buffer)
			if err != nil {
				break
			}
			if string(buffer[:bytesRead]) == "START" {
				// 更新文件
				common.Log.Info("Received start updating identifier")
				updateFile(conn)
			}
		}
		conn.Close()
	}
}
