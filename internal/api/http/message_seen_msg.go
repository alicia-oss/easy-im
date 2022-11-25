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

func SeenMsgHandler(ctx *gin.Context) {
	req, resp := &pb.SeenMsgReq{}, &pb.SeenMsgResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("SeenMsg bind err:%v", err), "api_http")
		return
	}

	resp = doSeenMsg(ctx, req)
	ctx.JSON(http.StatusOK, resp)
}

func doSeenMsg(ctx *gin.Context, req *pb.SeenMsgReq) (resp *pb.SeenMsgResp) {
	u, _ := ctx.Get(pkg.CTXUserId)
	uid := u.(uint64)
	err := domain.MessageService.UpdateMsgStateToSeen(req.GetMsgId(), uid, time.Now())
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
	_ = domain.ConnService.SendNotice(uid, int32(pb.NoticeType_message_seen), bytes)
	return resp
}
