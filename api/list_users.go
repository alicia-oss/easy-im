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

const ModuleNameListUsers = "handler_list_users"

type ListUsersHandler struct {
	jinx_imp.BaseMsgHandler
}

func (u *ListUsersHandler) Handle(request jinx_int.IRequest) {
	users := service.EasyIMUserManager.GetAllUser()
	var res []*pb.UserInfo
	for _, user := range users {
		t := &pb.UserInfo{
			UserId:   user.GetUserId(),
			UserName: user.GetUserName(),
		}
		res = append(res, t)
	}
	resp := pb.ListUsersResponse{UserInfos: res}
	bytes, _ := proto.Marshal(&resp)
	err := request.GetConnection().Send(bytes, uint32(pb.MessageType_LIST_USERS))
	if err != nil {
		log.Error(fmt.Sprintf("send error:%v", err), ModuleNameListUsers)
	}
}
