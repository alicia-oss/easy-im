package handler

import (
	"easy_im/pb"
	"easy_im/pkg/log"
	"fmt"
	client_imp "github.com/alicia-oss/jinx/client/imp"
	client_int "github.com/alicia-oss/jinx/client/int"
	"google.golang.org/protobuf/proto"
)

const ModuleNameReceiveMessage = "handler_receiver_message"

type ReceiveMessageHandler struct {
	client_imp.BaseMsgHandler
}

func (l *ReceiveMessageHandler) Handle(request client_int.IRequest) {
	message := &pb.Message{}
	err := proto.Unmarshal(request.GetData(), message)
	if err != nil {
		log.Error(fmt.Sprintf("Unmarshal error: %v", err), ModuleNameReceiveMessage)
		return
	}
	fmt.Printf("receive message :%v \n", message)
}
