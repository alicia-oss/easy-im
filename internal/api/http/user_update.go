package http

import (
	"easy_im/internal/api/pkg"
	userModel "easy_im/internal/domain/user/model"
	"easy_im/internal/domain/user/service"
	"easy_im/pb"
	"easy_im/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserUpdateInfoHandler(ctx *gin.Context) {
	req, resp := &pb.UpdateInfoReq{}, &pb.UpdateInfoResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("UpdateInfo bind err:%v", err), "api_http")
		return
	}

	resp = doUpdateInfo(ctx, req)
	ctx.JSON(http.StatusOK, resp)
}

func doUpdateInfo(ctx *gin.Context, req *pb.UpdateInfoReq) (resp *pb.UpdateInfoResp) {
	u := &userModel.User{
		ID:       req.UserId,
		Nickname: req.Nickname,
	}
	// todo gorm不完全更新操作验证
	err := service.UserService.Update(u)
	if err != nil {
		resp.Base = pkg.InternalError(err)
	} else {
		resp.Base = pkg.Success()
	}
	return resp
}

func UserUpdatePasswordHandler(ctx *gin.Context) {
	req, resp := &pb.UpdateUserPasswordReq{}, &pb.UpdateUserPasswordResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("UpdateUserPassword bind err:%v", err), "api_http")
		return
	}

	resp = doUpdateUserPassword(ctx, req)
	ctx.JSON(http.StatusOK, resp)
}

func doUpdateUserPassword(ctx *gin.Context, req *pb.UpdateUserPasswordReq) (resp *pb.UpdateUserPasswordResp) {
	u := &userModel.User{
		ID:       req.UserId,
		Password: req.Password,
	}
	err := service.UserService.Update(u)
	if err != nil {
		resp.Base = pkg.InternalError(err)
	} else {
		resp.Base = pkg.Success()
	}
	return resp
}
