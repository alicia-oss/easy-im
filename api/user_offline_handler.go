package api

import (
	"easy_im/pb"
	"easy_im/pkg/log"
	"easy_im/service"
	"fmt"
	"github.com/alicia-oss/jinx/jinx_imp"
	"github.com/alicia-oss/jinx/jinx_int"
	"google.golang.org/protobuf/proto"
)

const ModuleNameOffline = "handler_offline"

type UserOfflineHandler struct {
	jinx_imp.BaseMsgHandler
}

func (u *UserOfflineHandler) Handle(request jinx_int.IRequest) {
	bytes := request.GetData()
	req := &pb.OfflineReq{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		log.Error(fmt.Sprintf("UserOnlineHandler err:%v", err), ModuleNameOffline)
		return
	}
	service.EasyIMUserManager.RemoveUser(uint32(req.GetUserId()))
}
