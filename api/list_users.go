package api

import (
	"easy_im/pb"
	"easy_im/pkg/biz_coder"
	"easy_im/service"
	"github.com/alicia-oss/jinx/jinx_int"
	"google.golang.org/protobuf/proto"
)

const ModuleNameListUsers = "handler_list_users"

func NewListUsersHandler() jinx_int.IMsgHandle {
	return &TemplateHandler{
		BizCoder:   biz_coder.BizCoder{},
		BizHandler: listUsersHandler{},
	}
}

type listUsersHandler struct {
}

func (u listUsersHandler) HandleBiz(request jinx_int.IRequest, bytes []byte) (proto.Message, uint32, error) {
	users := service.EasyIMUserManager.GetAllUser()
	var res []*pb.UserInfo
	for _, user := range users {
		t := &pb.UserInfo{
			UserId:   user.GetUserId(),
			UserName: user.GetUserName(),
		}
		res = append(res, t)
	}
	resp := &pb.ListUsersResponse{UserInfos: res}
	return resp, uint32(pb.MessageType_LIST_USERS), nil
}
