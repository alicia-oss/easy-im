package http

import (
	"easy_im/internal/api/pkg"
	user_pkg "easy_im/internal/domain/user/pkg"
	"easy_im/internal/domain/user/service"
	"easy_im/pb"
	"easy_im/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthHandler(ctx *gin.Context, req *pb.AuthReq) *pb.AuthResp {
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
