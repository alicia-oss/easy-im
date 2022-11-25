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

func MessageGetSessionSeqHandler(ctx *gin.Context) {
	req, resp := &pb.GetSessionSeqReq{}, &pb.GetSessionSeqResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("GetSessionSeq bind err:%v", err), "api_http")
		return
	}

	resp = doGetSessionSeq(ctx, req)
	ctx.JSON(http.StatusOK, resp)
}

func doGetSessionSeq(ctx *gin.Context, req *pb.GetSessionSeqReq) (resp *pb.GetSessionSeqResp) {
	recvType, err := pkg.TransSessionTypeToRecvType(req.Type)
	if err != nil {
		resp.Base = pkg.UserError(err)
		return resp
	}
	seq, err := domain.UserSeqService.GenSeq(req.SessionId, recvType)
	if err != nil {
		resp.Base = pkg.InternalError(err)
	} else {
		resp.Seq = seq
	}
	return resp
}
