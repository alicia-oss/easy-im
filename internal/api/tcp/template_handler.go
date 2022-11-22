package tcp

import (
	"easy_im/internal/api/pkg"
	"easy_im/pb"
	"easy_im/pkg/biz_coder"
	"easy_im/pkg/jwt"
	"github.com/alicia-oss/jinx/jinx_int"
	"google.golang.org/protobuf/proto"
)

func NewBizHandler(f func(request jinx_int.IRequest, bytes []byte)) *TemplateHandler {
	handler := &TemplateHandler{
		BizCoder: biz_coder.BizCoder{},
		f:        f,
	}
	return handler
}

type TemplateHandler struct {
	biz_coder.BizCoder
	f func(request jinx_int.IRequest, bytes []byte)
}

func (u *TemplateHandler) Handle(ctx jinx_int.IRequest) {
	bytes := ctx.GetData()
	req, err := u.Decode(bytes)
	if err != nil {
		return
	}
	c, err := jwtAuth(req)
	if err != nil {
		ack := &pb.Ack{
			Type: 0,
			Id:   req.RequestId,
			Data: []byte("auth failed"),
		}
		bytes, _ := proto.Marshal(ack)
		_ = ctx.GetConnection().Send(bytes, 10000)
	}
	ctx.SetAttr(pkg.CTXUserId, c.UserId)
	ctx.SetAttr(pkg.CTXUserName, c.UserName)
	ctx.SetAttr(pkg.CTXRequestId, req.RequestId)
	u.f(ctx, req.GetData())
}

func (u *TemplateHandler) PreHandle(request jinx_int.IRequest) {
}

func (u *TemplateHandler) PostHandle(request jinx_int.IRequest) {
}

func jwtAuth(req *pb.Request) (*jwt.ImClaims, error) {
	token := req.GetToken()
	claims, err := jwt.DecodeToken(token)
	if err != nil {
		return claims, err
	}
	return claims, nil
}
