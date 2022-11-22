package tcp

import (
	"easy_im/internal/api/pkg"
	conn "easy_im/internal/domain/conn/service"
	"easy_im/pb"
	jinx "github.com/alicia-oss/jinx/jinx_int"
	"google.golang.org/protobuf/proto"
)

func OfflineHandler(ctx jinx.IRequest, bytes []byte) {
	var resp *pb.OfflineResp
	u, _ := ctx.GetAttr(pkg.CTXUserId)
	r, _ := ctx.GetAttr(pkg.CTXRequestId)
	uid, rid := u.(uint64), r.(string)

	err := conn.ConnService.OfflineUser(uid)
	if err != nil {
		resp = &pb.OfflineResp{
			Base: pkg.UserError(err),
		}
	} else {
		resp = &pb.OfflineResp{
			Base: pkg.Success(),
		}
	}

	bytes, _ = proto.Marshal(resp)
	_ = conn.ConnService.SendAck(uid, rid, bytes)

}
