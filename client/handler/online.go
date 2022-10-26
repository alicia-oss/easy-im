package handler

import (
	"easy_im/client/pkg"
	"easy_im/pb"
	"easy_im/pkg/log"
	"fmt"
	client_imp "github.com/alicia-oss/jinx/client/imp"
	client_int "github.com/alicia-oss/jinx/client/int"
	"google.golang.org/protobuf/proto"
)

const ModuleNameOnline = "handler_online"

type OnlineHandler struct {
	client_imp.BaseMsgHandler
}

func (l *OnlineHandler) Handle(request client_int.IRequest) {
	output, _ := pkg.BizCoderImpl.Decode(request.GetData())
	message := &pb.OnlineResponse{}
	err := proto.Unmarshal(output.GetData(), message)
	if err != nil {
		log.Error(fmt.Sprintf("Unmarshal error: %v", err), ModuleNameOnline)
		return
	}
	fmt.Printf("online... your user_id:%v, user_name:%v \n", message.GetUserId(), message.GetUserName())
	request.GetClient().SetAttr(pkg.AttrUserId, message.GetUserId())
	request.GetClient().SetAttr(pkg.AttrUserName, message.GetUserName())
}
