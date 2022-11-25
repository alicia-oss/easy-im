package http

import (
	"easy_im/internal/api/pkg"
	"easy_im/internal/domain"
	"easy_im/pb"
	"easy_im/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SyncSessionHandler(ctx *gin.Context) {
	req, resp := &pb.SyncSessionReq{}, &pb.SyncSessionResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("SyncSession bind err:%v", err), "api_http")
		return
	}
	resp = doSyncSession(ctx, req)
	ctx.JSON(http.StatusOK, resp)
}

func doSyncSession(ctx *gin.Context, req *pb.SyncSessionReq) (resp *pb.SyncSessionResp) {
	u, _ := ctx.Get(pkg.CTXUserId)
	uid := u.(uint64)

	messages, err := domain.MessageService.SyncInbox(uid, req.SessionId, req.MinSeq, req.MaxSeq)
	if err != nil {
		resp.Base = pkg.InternalError(err)
	} else {
		resp.Base = pkg.Success()
		list := make([]*pb.Message, 0, len(messages))
		for _, msg := range messages {
			vo := pkg.TransMsgDoToVo(msg)
			if vo == nil {
				continue
			}
			list = append(list, vo)
		}
		resp.Messages = list
	}
	return resp
}
