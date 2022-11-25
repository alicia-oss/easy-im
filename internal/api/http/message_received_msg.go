package http

import (
	"easy_im/internal/api/pkg"
	"easy_im/internal/domain"
	pkg2 "easy_im/internal/domain/message/pkg"
	"easy_im/pb"
	"easy_im/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
	"net/http"
	"time"
)

func ReceivedMsgHandler(ctx *gin.Context) {
	req, resp := &pb.ReceivedMsgReq{}, &pb.ReceivedMsgResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("ReceivedMsg bind err:%v", err), "api_http")
		return
	}

	resp = doReceivedMsg(ctx, req)
	ctx.JSON(http.StatusOK, resp)
}

func doReceivedMsg(ctx *gin.Context, req *pb.ReceivedMsgReq) (resp *pb.ReceivedMsgResp) {
	u, _ := ctx.Get(pkg.CTXUserId)
	uid := u.(uint64)
	err := domain.MessageService.UpdateMsgStateToDelivered(req.GetMsgId(), uid, time.Now())
	if err != nil {
		switch err {
		case pkg2.InvalidUserId:
			resp.Base = pkg.UserError(err)
		case pkg2.ErrUnknown:
			resp.Base = pkg.InternalError(err)
		}
		return resp
	}
	// notice user
	notice := &pb.MsgStateChangeNotice{
		MsgId: req.GetMsgId(),
		State: pb.MessageState_Seen,
	}
	bytes, _ := proto.Marshal(notice)
	_ = domain.ConnService.SendNotice(uid, int32(pb.NoticeType_message_delivered), bytes)
	return resp
}
