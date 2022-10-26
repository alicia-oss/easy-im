package handler

import (
	"easy_im/client/pkg"
	"easy_im/pb"
	"easy_im/pkg/log"
	"fmt"
	client_imp "github.com/alicia-oss/jinx/client/imp"
	"github.com/alicia-oss/jinx/client/int"
	"google.golang.org/protobuf/proto"
)

const ModuleNameListUsers = "handler_list_users"

type ListUsersHandler struct {
	client_imp.BaseMsgHandler
}

func (l *ListUsersHandler) Handle(request client_int.IRequest) {
	output, _ := pkg.BizCoderImpl.Decode(request.GetData())
	resp := &pb.ListUsersResponse{}
	err := proto.Unmarshal(output.GetData(), resp)
	if err != nil {
		log.Error(fmt.Sprintf("Unmarshal error: %v", err), ModuleNameListUsers)
		return
	}
	fmt.Printf("user list: %v \n", resp.UserInfos)
}
