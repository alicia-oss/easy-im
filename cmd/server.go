package main

import (
	"easy_im/internal/api/http"
	"github.com/gin-gonic/gin"
)

func main() {
	//server := jinx_imp.NewServer("EasyIM", "tcp", "127.0.0.1", 9990)
	//server.AddRouter(uint32(pb.RequestType_ONLINE), api.NewUserOnlineHandler())
	//server.AddRouter(uint32(pb.MessageType_OFFLINE), api.NewUserOfflineHandler())
	//server.AddRouter(uint32(pb.MessageType_SEND_MESSAGE), api.NewSendMessageHandler())
	//server.AddRouter(uint32(pb.MessageType_LIST_USERS), api.NewListUsersHandler())
	//server.SetOnCloseHandler(&api.OnConnCloseHandler{})
	//server.Start()

	r := gin.Default()
	r.POST("/auth", http.AuthHandler)
	r.Run("127.0.0.1:8080")

}
