# 脚本管理服务主机终端（client）

## install

```bash
go mod tidy
go build -o scriptsclient main.go
```

## arg

```bash
--address          # 服务端地址
--reconnect        # 重新连接服务端重试间隔，默认0不重试
--script           # 脚本文件输出的文件路径，string类型的数组，目前一个客户端仅支持维护一个脚本，所以会数组第一个值
```

## start up 

```bash
nohup ./scriptsclient --address=127.0.0.1:7070 --reconnect=10  --script=shell.sh > scriptsclient.log 2>&1 &

```