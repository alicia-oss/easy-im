package api

import (
	"easy_im/service"
	"github.com/alicia-oss/jinx/jinx_int"
)

type OnConnCloseHandler struct{}

func (o *OnConnCloseHandler) Handle(conn jinx_int.IConnection) {
	attr, b := conn.GetAttr("user_id")
	if b {
		uid := attr.(uint32)
		service.EasyIMUserManager.RemoveUser(uid)
	}
}
