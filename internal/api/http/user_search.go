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

func SearchUserHandler(ctx *gin.Context) {
	req, resp := &pb.SearchUserReq{}, &pb.SearchUserResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("SearchUser bind err:%v", err), "api_http")
		return
	}

	resp = doSearchUser(ctx, req)
	ctx.JSON(http.StatusOK, resp)
}

func doSearchUser(ctx *gin.Context, req *pb.SearchUserReq) (resp *pb.SearchUserResp) {
	users, err := service.UserService.Search(req.GetKey())
	if err != nil {
		resp.Base = pkg.InternalError(err)
	} else {
		resp.Base = pkg.Success()
		resp.Users = userMoToVo(users)
	}
	return resp
}
