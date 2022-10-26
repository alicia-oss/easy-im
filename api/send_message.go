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

const ModuleNameSendMessage = "handler_send_message"

type SendMessageHandler struct {
	jinx_imp.BaseMsgHandler
}

func (u *SendMessageHandler) Handle(request jinx_int.IRequest) {
	bytes := request.GetData()
	req := &pb.SendMessageReq{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		log.Error(fmt.Sprintf("UserOnlineHandler err:%v", err), ModuleNameSendMessage)
		return
	}
	sender, ok := service.EasyIMUserManager.GetUserById(req.GetSenderId())
	if !ok {
		return
	}
	receiver, ok := service.EasyIMUserManager.GetUserById(req.GetReceiverId())
	if !ok {
		return
	}
	message := pb.Message{
		SenderId:     req.SenderId,
		SenderName:   sender.GetUserName(),
		ReceiverId:   req.GetReceiverId(),
		ReceiverName: receiver.GetUserName(),
		Data:         req.GetMessageContent(),
	}
	bytes, _ = proto.Marshal(&message)
	receiver.SendMessage(bytes, uint32(pb.MessageType_RECEIVE_MESSAGE))

}
