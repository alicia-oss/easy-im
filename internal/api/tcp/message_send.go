package tcp

import (
	"easy_im/internal/api/pkg"
	"easy_im/internal/domain"
	"easy_im/internal/domain/message/model"
	pkg2 "easy_im/internal/domain/message/pkg"
	"easy_im/pb"
	jinx "github.com/alicia-oss/jinx/jinx_int"
	"google.golang.org/protobuf/proto"
	"time"
)

func MessageSendHandler(ctx jinx.IRequest, bytes []byte) {
	u, _ := ctx.GetAttr(pkg.CTXUserId)
	r, _ := ctx.GetAttr(pkg.CTXRequestId)
	uid, rid := u.(uint64), r.(string)
	req, resp, message, err := doSaveMessage(bytes, uid)
	// 2. ack
	bytes, _ = proto.Marshal(resp)
	_ = domain.ConnService.SendAck(uid, rid, bytes)

	if err != nil {
		return
	}
	// 3. if B online, deliver message
	switch message.ReceiverType {
	case pkg2.ReceiverType_USER:
		if domain.ConnService.GetUserState(req.ReceiverId) {
			doNoticeMsgToUser(message)
		}
	case pkg2.ReceiverType_GROUP:
		//TODO
	}

}

func doNoticeMsgToUser(message *model.Message) bool {
	if !domain.ConnService.GetUserState(message.ReceiverId) {
		return false
	}
	vo := pkg.TransMsgDoToVo(message)
	if vo == nil {
		return false
	}
	notice := &pb.NewMessageNotice{
		Msg: vo,
	}
	bytes, _ := proto.Marshal(notice)
	_ = domain.ConnService.SendNotice(message.ReceiverId, int32(pb.NoticeType_message), bytes)
	return true
}

func doSaveMessage(bytes []byte, uid uint64) (*pb.SendMessageReq, *pb.SendMessageResp, *model.Message, error) {
	var req *pb.SendMessageReq
	var resp = &pb.SendMessageResp{}
	err := proto.Unmarshal(bytes, req)
	if err != nil {
		resp.Base = pkg.UserError(err)
		return req, resp, nil, err
	}
	recvType, err := pkg.TransSessionTypeToRecvType(req.SessionType)
	if err != nil {
		resp.Base = pkg.UserError(err)
		return req, resp, nil, err
	}
	mType, err := pkg.TransMessageType(req.Type)
	if err != nil {
		resp = &pb.SendMessageResp{Base: pkg.UserError(err)}
		return req, resp, nil, err
	}
	// 1.save message
	message := &model.Message{
		SenderType:   pkg2.SenderType_USER,
		SenderId:     uid,
		ReceiverType: recvType,
		ReceiverId:   req.ReceiverId,
		Type:         mType,
		Content:      req.GetMessageContent(),
		Seq:          req.GetSeq(),
		State:        pkg2.MessageState_SENT,
		SentTime:     time.Time{},
		DeliverTime:  time.Time{},
		SeenTime:     time.Time{},
	}
	err = domain.MessageService.SaveMessage(message)
	if err != nil {
		resp.Base = pkg.InternalError(err)
		return req, resp, nil, err
	}
	resp.Base = pkg.Success()
	return req, resp, message, nil
}
