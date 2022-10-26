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

type SendMessageHandler struct {
	client_imp.BaseMsgHandler
}

func (l *SendMessageHandler) Handle(request client_int.IRequest) {
	output, _ := pkg.BizCoderImpl.Decode(request.GetData())
	message := &pb.SendMessageResponse{}
	err := proto.Unmarshal(output.GetData(), message)
	if err != nil {
		log.Error(fmt.Sprintf("Unmarshal error: %v", err), ModuleNameReceiveMessage)
		return
	}
	fmt.Printf("resp message :%v \n", output)
}
