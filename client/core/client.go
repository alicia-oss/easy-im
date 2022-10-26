package core

import (
	"easy_im/client/handler"
	"easy_im/client/pkg"
	"easy_im/pb"
	client_imp "github.com/alicia-oss/jinx/client/imp"
	client_int "github.com/alicia-oss/jinx/client/int"
	"google.golang.org/protobuf/proto"
)

func NewClient(userName string) (IClient, error) {
	conn, err := client_imp.NewClient("tcp", "127.0.0.1", 9990)
	if err != nil {
		return nil, err
	}
	c := &client{
		IClient: conn,
	}

	c.AddRoute(uint32(pb.MessageType_LIST_USERS), &handler.ListUsersHandler{})
	c.AddRoute(uint32(pb.MessageType_RECEIVE_MESSAGE), &handler.ReceiveMessageHandler{})
	c.AddRoute(uint32(pb.MessageType_ONLINE), &handler.OnlineHandler{})
	c.Start()
	c.Online(userName)
	return c, nil
}

type client struct {
	client_int.IClient
}

func (c *client) Online(userName string) {
	req := &pb.OnlineReq{UserName: userName}
	bytes, _ := proto.Marshal(req)
	c.Send(bytes, uint32(pb.MessageType_ONLINE))
}

func (c *client) Offline() {
	req := &pb.OfflineReq{}
	bytes, _ := proto.Marshal(req)
	c.Send(bytes, uint32(pb.MessageType_OFFLINE))
	c.Close()
}

func (c *client) ListUsers() {
	req := &pb.ListUsersReq{}
	bytes, _ := proto.Marshal(req)
	c.Send(bytes, uint32(pb.MessageType_LIST_USERS))
}

func (c *client) SendMessage(content string, receiverId uint32) {
	userId, _ := c.GetUserId()
	req := &pb.SendMessageReq{
		SenderId:       userId,
		ReceiverId:     receiverId,
		MessageContent: content,
	}
	bytes, _ := proto.Marshal(req)
	c.Send(bytes, uint32(pb.MessageType_SEND_MESSAGE))
}

func (c *client) GetUserId() (uint32, bool) {
	attr, b := c.GetAttr(pkg.AttrUserId)
	if b {
		return attr.(uint32), b
	}
	return 0, b
}
func (c *client) GetUserName() (string, bool) {
	attr, b := c.GetAttr(pkg.AttrUserName)
	if b {
		return attr.(string), b
	}
	return "", b
}
func (c *client) SetUserId(uId uint32) {
	c.SetAttr(pkg.AttrUserId, uId)
}
func (c *client) SetUserName(name string) {
	c.SetAttr(pkg.AttrUserName, name)
}
