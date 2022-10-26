package main

import (
	"easy_im/api"
	"easy_im/pb"
	"github.com/alicia-oss/jinx/jinx_imp"
)

func main() {
	server := jinx_imp.NewServer("EasyIM", "tcp", "127.0.0.1", 9990)
	server.AddRouter(uint32(pb.MessageType_ONLINE), &api.UserOnlineHandler{})
	server.AddRouter(uint32(pb.MessageType_OFFLINE), &api.UserOfflineHandler{})
	server.AddRouter(uint32(pb.MessageType_SEND_MESSAGE), &api.SendMessageHandler{})
	server.AddRouter(uint32(pb.MessageType_LIST_USERS), &api.ListUsersHandler{})
	server.SetOnCloseHandler(&api.OnConnCloseHandler{})
	server.Start()

	q := make(chan struct{})
	<-q
	server.Stop()
}
