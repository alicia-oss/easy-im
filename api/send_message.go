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

const ModuleNameSendMessage = "handler_send_message"

func NewSendMessageHandler() jinx_int.IMsgHandle {
	return &TemplateHandler{
		BizCoder:   biz_coder.BizCoder{},
		BizHandler: sendMessageHandler{},
	}
}

type sendMessageHandler struct {
}

func (u sendMessageHandler) HandleBiz(request jinx_int.IRequest, data []byte) (proto.Message, uint32, error) {
	req := &pb.SendMessageReq{}
	if err := proto.Unmarshal(data, req); err != nil {
		log.Error(fmt.Sprintf("UserOnlineHandler err:%v", err), ModuleNameSendMessage)
		return nil, uint32(pb.MessageType_SEND_MESSAGE), errors.NewInternalError()
	}
	sender, ok := service.EasyIMUserManager.GetUserById(req.GetSenderId())
	if !ok {
		return nil, uint32(pb.MessageType_SEND_MESSAGE), errors.NewInternalError()
	}
	receiver, ok := service.EasyIMUserManager.GetUserById(req.GetReceiverId())
	if !ok {
		return nil, uint32(pb.MessageType_SEND_MESSAGE), errors.NewInternalError()
	}
	message := pb.Message{
		SenderId:     req.SenderId,
		SenderName:   sender.GetUserName(),
		ReceiverId:   req.GetReceiverId(),
		ReceiverName: receiver.GetUserName(),
		Data:         req.GetMessageContent(),
	}
	bytes, _ := proto.Marshal(&message)
	receiver.SendMessage(bytes, uint32(pb.MessageType_RECEIVE_MESSAGE))
	return &pb.SendMessageResponse{}, uint32(pb.MessageType_SEND_MESSAGE), nil
}
