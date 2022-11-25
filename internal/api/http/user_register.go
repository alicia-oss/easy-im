package http

import (
	"easy_im/internal/api/pkg"
	"easy_im/internal/domain"
	userPkg "easy_im/internal/domain/user/pkg"
	"easy_im/pb"
	"easy_im/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandler(ctx *gin.Context) {
	req, resp := &pb.RegisterReq{}, &pb.RegisterResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("Register bind err:%v", err), "api_http")
		return
	}

	resp = doRegister(ctx, req)
	ctx.JSON(http.StatusOK, resp)
}

//TODO Transaction support
func doRegister(ctx *gin.Context, req *pb.RegisterReq) (resp *pb.RegisterResp) {
	u, err := domain.UserService.Register(req.Username, req.GetPassword(), req.GetPassword())
	if err == userPkg.ErrUsernameUsed {
		resp.Base = pkg.UserError(err)
		return resp
	} else if err == userPkg.ErrUnknown {
		resp.Base = pkg.InternalError(err)
		return resp
	}
	err = domain.UserSeqService.CreateSeqBox(u.ID)
	if err != nil {
		resp.Base = pkg.InternalError(err)
	} else {
		resp.Base = pkg.Success()
	}
	return resp
}
