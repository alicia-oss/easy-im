package api

import (
	"easy_im/pb"
	"easy_im/pkg/log"
	"easy_im/service"
	user_imp "easy_im/service/user_manager/imp"
	"fmt"
	"github.com/alicia-oss/jinx/jinx_imp"
	"github.com/alicia-oss/jinx/jinx_int"
	"google.golang.org/protobuf/proto"
)

const ModuleNameOnline = "handler_online"

type UserOnlineHandler struct {
	jinx_imp.BaseMsgHandler
}

func (u *UserOnlineHandler) Handle(request jinx_int.IRequest) {
	bytes := request.GetData()
	req := &pb.OnlineReq{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		log.Error(fmt.Sprintf("UserOnlineHandler err:%v", err), ModuleNameOnline)
		return
	}
	user := user_imp.NewUser(req.GetUserName(), request.GetConnection(), service.EasyIMUserManager)
	service.EasyIMUserManager.AddUser(user)
	request.GetConnection().SetAttr("user_id", user.GetUserId())

	resp := &pb.OnlineResponse{
		UserId:   user.GetUserId(),
		UserName: user.GetUserName(),
	}
	bytes, err := proto.Marshal(resp)
	if err != nil {
		log.Error(fmt.Sprintf("Marshal err:%v", err), ModuleNameOnline)
		return
	}
	request.GetConnection().Send(bytes, uint32(pb.MessageType_ONLINE))
}
