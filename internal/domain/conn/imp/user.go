package user_imp

import (
	"easy_im/service/user_manager/int"
	"github.com/alicia-oss/jinx/jinx_int"
	"math/rand"
)

func NewUser(userName string, conn jinx_int.IConnection, userManager user_int.IUserManager) user_int.IUser {
	userId := rand.Uint64()
	return &user{userId: uint(userId), userName: userName, conn: conn, userManager: userManager}
}

type user struct {
	userId      uint
	userName    string
	conn        jinx_int.IConnection
	userManager user_int.IUserManager
}

func (u *user) GetUserId() uint {
	return u.userId
}

func (u *user) GetUserName() string {
	return u.userName
}

func (u *user) GetUserIP() string {
	return u.conn.GetRemoteAddr().String()
}

func (u *user) Offline() {
	u.userManager.RemoveUser(uint32(u.userId))
}

func (u *user) Online() {
	u.userManager.AddUser(u)
}

func (u *user) SendMessage(data []byte, msgType uint32) error {
	if err := u.conn.Send(data, msgType); err != nil {
		return err
	}
	return nil
}
