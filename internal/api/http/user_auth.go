package http

import (
	"easy_im/internal/api/pkg"
	user_pkg "easy_im/internal/domain/user/pkg"
	"easy_im/internal/domain/user/service"
	"easy_im/pb"
	"easy_im/pkg/jwt"
	"easy_im/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthHandler(ctx *gin.Context) {
	req, resp := &pb.AuthReq{}, &pb.AuthResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("auth bind err:%v", err), "api_http")
		return
	}

	resp = doAuthHandler(err, req)
	ctx.JSON(http.StatusOK, resp)
}

func doAuthHandler(err error, req *pb.AuthReq) (resp *pb.AuthResp) {
	err, user := service.UserService.Auth(req.GetUsername(), req.GetPassword())
	if err != nil {
		switch err {
		case user_pkg.ErrUnknown:
			return &pb.AuthResp{
				Base:  pkg.InternalError(err),
				Token: "",
			}
		default:
			return &pb.AuthResp{
				Base:  pkg.UserError(err),
				Token: "",
			}
		}
	}

	token, _ := jwt.GenToken(user.ID, user.Username)
	return &pb.AuthResp{
		Base:  pkg.Success(),
		Token: token,
	}
}
