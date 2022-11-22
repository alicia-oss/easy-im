package tcp

import (
	"easy_im/internal/api/pkg"
	conn "easy_im/internal/domain/conn/service"
	user "easy_im/internal/domain/user/service"
	"easy_im/pb"
	jinx "github.com/alicia-oss/jinx/jinx_int"
	"google.golang.org/protobuf/proto"
)

func LoginHandler(ctx jinx.IRequest, bytes []byte) {
	jinxConn := ctx.GetConnection()
	u, _ := ctx.GetAttr(pkg.CTXUserId)
	r, _ := ctx.GetAttr(pkg.CTXRequestId)
	uid, rid := u.(uint64), r.(string)
	conn.ConnService.OnlineUser(uid, jinxConn)
	userT, _ := user.UserService.GetById(uid)
	resp := &pb.OnlineResp{Base: pkg.Success(), Info: &pb.User{
		Id:       userT.ID,
		Username: userT.Username,
		Nickname: userT.Nickname,
		State:    pb.UserState_ON,
	}}

	bytes, _ = proto.Marshal(resp)
	_ = conn.ConnService.SendAck(uid, rid, bytes)
}
