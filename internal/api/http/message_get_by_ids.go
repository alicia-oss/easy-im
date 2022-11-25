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

func GetMessagesHandler(ctx *gin.Context) {
	req, resp := &pb.GetMessagesReq{}, &pb.GetMessagesResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("GetMessages bind err:%v", err), "api_http")
		return
	}

	resp = doGetMessages(ctx, req)
	ctx.JSON(http.StatusOK, resp)
}

func doGetMessages(ctx *gin.Context, req *pb.GetMessagesReq) (resp *pb.GetMessagesResp) {
	doMsgs, err := domain.MessageService.GetMessageByIds(req.GetMsgIds())
	if err != nil {
		resp.Base = pkg.InternalError(err)
	} else {
		list := make([]*pb.Message, len(doMsgs))
		for i, msg := range doMsgs {
			list[i] = pkg.TransMsgDoToVo(msg)
		}
		resp.Messages = list
		resp.Base = pkg.Success()
	}
	return resp
}
