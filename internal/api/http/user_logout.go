package http

import (
	"easy_im/internal/api/pkg"
	"easy_im/internal/domain/user/service"
	"easy_im/pb"
	"easy_im/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutHandler(ctx *gin.Context) {
	req, resp := &pb.LogoutReq{}, &pb.LogoutResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("Logout bind err:%v", err), "api_http")
		return
	}

	resp = doLogout(ctx, req)
	ctx.JSON(http.StatusOK, resp)
}

func doLogout(ctx *gin.Context, req *pb.LogoutReq) (resp *pb.LogoutResp) {
	u, _ := ctx.Get(pkg.CTXUserId)
	uid := u.(uint64)
	err := service.UserService.Logout(uid)
	if err != nil {
		resp.Base = pkg.InternalError(err)
	} else {
		resp.Base = pkg.Success()
	}
	return resp
}
