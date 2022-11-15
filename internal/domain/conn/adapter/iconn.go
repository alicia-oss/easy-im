package adapter

import "net"

type IConnection interface {
	// Stop 停止链接， 结束当前链接
	Stop()
	// GetRemoteAddr 获取远程客户端的TCP状态 IP Port
	GetRemoteAddr() net.Addr
	// Send 发送数据，将数据发送给client
	Send(data []byte, msgId uint32) error
}
