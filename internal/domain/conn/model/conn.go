package model

import (
	"easy_im/internal/domain/conn/adapter"
	"easy_im/internal/domain/conn/repo"
	"easy_im/pb"
	"google.golang.org/protobuf/proto"
)

type IConn interface {
	GetUserId() uint64
	GetUserIP() string
	Offline()
	Online()
	SendNotice(data []byte, nType int32, noticeId string) error
	SendAck(data []byte, id string) error
}

func NewConn(userId uint64, conn adapter.IConnection, userManager repo.IConnManager) IConn {
	return &connImpl{userId: userId, conn: conn, userManager: userManager}
}

type connImpl struct {
	userId      uint64
	conn        adapter.IConnection
	userManager repo.IConnManager
}

func (u *connImpl) SendNotice(data []byte, nType int32, noticeId string) error {
	n := &pb.Notice{
		NoticeId: noticeId,
		Data:     data,
		Type:     pb.NoticeType(nType),
	}
	bytes, _ := proto.Marshal(n)
	return u.sendMessage(bytes, nType)
}

func (u *connImpl) SendAck(data []byte, id string) error {
	ack := &pb.Ack{
		Type: pb.AckType_REQUEST,
		Id:   id,
		Data: data,
	}
	bytes, _ := proto.Marshal(ack)
	return u.sendMessage(bytes, 1000)
}

func (u *connImpl) GetUserId() uint64 {
	return u.userId
}

func (u *connImpl) GetUserIP() string {
	return u.conn.GetRemoteAddr().String()
}

func (u *connImpl) Offline() {
	u.userManager.RemoveConn(u.userId)
	u.conn.Stop()
}

func (u *connImpl) Online() {
	u.userManager.AddConn(u)
}

func (u *connImpl) sendMessage(data []byte, msgType int32) error {
	if err := u.conn.Send(data, uint32(msgType)); err != nil {
		return err
	}
	return nil
}
