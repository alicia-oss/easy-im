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

func MessageSyncInboxHandler(ctx *gin.Context) {
	req, resp := &pb.SyncInboxReq{}, &pb.SyncInboxResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("SyncInbox bind err:%v", err), "api_http")
		return
	}
	resp = doSyncInbox(ctx, req)
	ctx.JSON(http.StatusOK, resp)
}

func doSyncInbox(ctx *gin.Context, req *pb.SyncInboxReq) (resp *pb.SyncInboxResp) {
	u, _ := ctx.Get(pkg.CTXUserId)
	uid := u.(uint64)
	ids, err := domain.MessageService.GetMessageIdsByUserId(uid)
	if err != nil {
		resp.Base = pkg.InternalError(err)
	} else {
		resp.Base = pkg.Success()
		resp.MsgIds = ids
	}
	return resp
}
