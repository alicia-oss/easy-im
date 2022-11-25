package pkg

import (
	"easy_im/internal/domain/message/model"
	"easy_im/internal/domain/message/pkg"
	"easy_im/pb"
	"easy_im/pkg/log"
	"errors"
	"fmt"
)

var (
	receiverTypeMap = map[pb.SessionType]int8{
		pb.SessionType_USER:  pkg.ReceiverType_USER,
		pb.SessionType_GROUP: pkg.ReceiverType_GROUP,
	}

	sessionTypeMap = map[int8]pb.SessionType{
		pkg.ReceiverType_USER:  pb.SessionType_USER,
		pkg.ReceiverType_GROUP: pb.SessionType_GROUP,
	}

	messageTypeMap = map[pb.MessageType]int8{
		pb.MessageType_IMG: pkg.MessageType_IMG,
		pb.MessageType_TXT: pkg.MessageType_TXT,
	}

	doMessageTypeMap = map[int8]pb.MessageType{
		pkg.MessageType_IMG: pb.MessageType_IMG,
		pkg.MessageType_TXT: pb.MessageType_TXT,
	}

	messageTypeMapToVo = map[int8]pb.MessageType{
		pkg.MessageType_IMG: pb.MessageType_IMG,
		pkg.MessageType_TXT: pb.MessageType_TXT,
	}
)

func TransSessionTypeToRecvType(sessionType pb.SessionType) (int8, error) {
	v, ok := receiverTypeMap[sessionType]
	if !ok {
		err := errors.New(fmt.Sprintf("invalid sessionType:%v", sessionType))
		log.Error(fmt.Sprintf("transSessionTypeToRecvType err:%v", err.Error()), "api_transport")
		return 0, err
	}
	return v, nil
}

func TransRecvTypeToSessionType(t int8) (pb.SessionType, error) {
	v, ok := sessionTypeMap[t]
	if !ok {
		err := errors.New(fmt.Sprintf("invalid sessionType:%v", t))
		log.Error(fmt.Sprintf("transSessionTypeToRecvType err:%v", err.Error()), "api_transport")
		return 0, err
	}
	return v, nil
}

func TransMessageType(t pb.MessageType) (int8, error) {
	v, ok := messageTypeMap[t]
	if !ok {
		err := errors.New(fmt.Sprintf("invalid message type:%v", t))
		log.Error(fmt.Sprintf("TransMessageType err:%v", err.Error()), "api_transport")
		return 0, err
	}
	return v, nil
}

func TransVoMessageType(t int8) (pb.MessageType, error) {
	v, ok := doMessageTypeMap[t]
	if !ok {
		err := errors.New(fmt.Sprintf("invalid message type:%v", t))
		log.Error(fmt.Sprintf("TransMessageType err:%v", err.Error()), "api_transport")
		return 0, err
	}
	return v, nil
}

func TransMessageTypeToVo(t int8) (pb.MessageType, error) {
	v, ok := messageTypeMapToVo[t]
	if !ok {
		err := errors.New(fmt.Sprintf("invalid message type:%v", t))
		log.Error(fmt.Sprintf("TransMessageTypeToVo err:%v", err.Error()), "api_transport")
		return 0, err
	}
	return v, nil
}

func TransMsgDoToVo(message *model.Message) *pb.Message {
	sessionType, err := TransRecvTypeToSessionType(message.ReceiverType)
	messageType, err := TransVoMessageType(message.Type)
	if err != nil {
		log.Error(fmt.Sprintf("TransMsgDoToVo err, msg:%v", message), "api_transport")
		return nil
	}
	p := &pb.Message{
		MsgId:       message.ID,
		SessionType: sessionType,
		SessionId:   message.ReceiverId,
		SenderId:    message.SenderId,
		Seq:         message.Seq,
		State:       pb.MessageState_Sent,
		SentTime:    message.SentTime.Format("2006-01-02 15:04:05"),
		DeliverTime: message.DeliverTime.Format("2006-01-02 15:04:05"),
		SeenTime:    message.SentTime.Format("2006-01-02 15:04:05"),
		Type:        messageType,
		Data:        message.Content,
	}
	return p
}
