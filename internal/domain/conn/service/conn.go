package service

import (
	"github.com/alicia-oss/jinx/jinx_int"
)

func NewUser(userId uint64, conn jinx_int.IConnection, userManager IConnManager) IConn {
	return &connImpl{userId: userId, conn: conn, userManager: userManager}
}

type connImpl struct {
	userId      uint64
	conn        jinx_int.IConnection
	userManager IConnManager
}

func (u *connImpl) GetUserId() uint64 {
	return u.userId
}

func (u *connImpl) GetUserIP() string {
	return u.conn.GetRemoteAddr().String()
}

func (u *connImpl) Offline() {
	u.userManager.RemoveConn(u.userId)
}

func (u *connImpl) Online() {
	u.userManager.AddConn(u)
}

func (u *connImpl) SendMessage(data []byte, msgType uint32) error {
	if err := u.conn.Send(data, msgType); err != nil {
		return err
	}
	return nil
}
