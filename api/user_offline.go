package api

import (
	"easy_im/pb"
	"easy_im/pkg/biz_coder"
	"easy_im/pkg/errors"
	"easy_im/pkg/log"
	"easy_im/service"
	"fmt"
	"github.com/alicia-oss/jinx/jinx_int"
	"google.golang.org/protobuf/proto"
)

const ModuleNameOffline = "handler_offline"

func NewUserOfflineHandler() jinx_int.IMsgHandle {
	return &TemplateHandler{
		BizCoder:   biz_coder.BizCoder{},
		BizHandler: userOfflineHandler{},
	}
}

type userOfflineHandler struct {
}

func (u userOfflineHandler) HandleBiz(request jinx_int.IRequest, bytes []byte) (proto.Message, uint32, error) {
	req := &pb.OfflineReq{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		log.Error(fmt.Sprintf("UserOnlineHandler err:%v", err), ModuleNameOffline)
		return nil, uint32(pb.MessageType_OFFLINE), errors.NewInternalError()
	}
	service.EasyIMUserManager.RemoveUser(uint32(req.GetUserId()))
	return &pb.OfflineResponse{}, uint32(pb.MessageType_OFFLINE), nil
}
