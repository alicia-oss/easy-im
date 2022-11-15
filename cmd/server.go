package main

import (
	"easy_im/api"
	"easy_im/pb"
	"github.com/alicia-oss/jinx/jinx_imp"
)

func main() {
	server := jinx_imp.NewServer("EasyIM", "tcp", "127.0.0.1", 9990)
	server.AddRouter(uint32(pb.RequestType_ONLINE), api.NewUserOnlineHandler())
	server.AddRouter(uint32(pb.MessageType_OFFLINE), api.NewUserOfflineHandler())
	server.AddRouter(uint32(pb.MessageType_SEND_MESSAGE), api.NewSendMessageHandler())
	server.AddRouter(uint32(pb.MessageType_LIST_USERS), api.NewListUsersHandler())
	server.SetOnCloseHandler(&api.OnConnCloseHandler{})
	server.Start()

}
