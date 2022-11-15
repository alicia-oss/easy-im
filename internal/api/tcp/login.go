package tcp

import (
	"easy_im/internal/domain/conn/service"
	"easy_im/pb"
	jinx "github.com/alicia-oss/jinx/jinx_int"
)

func LoginHandler(ctx jinx.IRequest, req *pb.OnlineReq) *pb.OnlineResp {
	jinxConn := ctx.GetConnection()
	ctx.GetAttr("")
	service.ConnService.OnlineUser()
}
