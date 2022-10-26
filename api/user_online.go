package api

import (
	"easy_im/pb"
	"easy_im/pkg/biz_coder"
	"easy_im/pkg/errors"
	"easy_im/pkg/log"
	"easy_im/service"
	user_imp "easy_im/service/user_manager/imp"
	"fmt"
	"github.com/alicia-oss/jinx/jinx_int"
	"google.golang.org/protobuf/proto"
)

const ModuleNameOnline = "handler_online"

func NewUserOnlineHandler() jinx_int.IMsgHandle {
	return &TemplateHandler{
		BizCoder:   biz_coder.BizCoder{},
		BizHandler: userOnlineHandler{},
	}
}

type userOnlineHandler struct {
}

func (u userOnlineHandler) HandleBiz(request jinx_int.IRequest, bytes []byte) (proto.Message, uint32, error) {
	req := &pb.OnlineReq{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		log.Error(fmt.Sprintf("UserOnlineHandler err:%v", err), ModuleNameOnline)
		return nil, uint32(pb.MessageType_ONLINE), errors.NewInternalError()
	}
	user := user_imp.NewUser(req.GetUserName(), request.GetConnection(), service.EasyIMUserManager)
	service.EasyIMUserManager.AddUser(user)
	request.GetConnection().SetAttr("user_id", user.GetUserId())

	resp := &pb.OnlineResponse{
		UserId:   user.GetUserId(),
		UserName: user.GetUserName(),
	}
	return resp, uint32(pb.MessageType_ONLINE), nil

}
