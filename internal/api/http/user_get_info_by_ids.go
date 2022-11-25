package http

import (
	"easy_im/internal/api/pkg"
	"easy_im/internal/domain"
	"easy_im/internal/domain/user/model"
	"easy_im/pb"
	"easy_im/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfoByIds(ctx *gin.Context) {
	req, resp := &pb.GetUserInfoByIdsReq{}, &pb.GetUserInfoByIdsResp{}
	err := ctx.Bind(req)
	if err != nil {
		log.Error(fmt.Sprintf("auth bind err:%v", err), "api_http")
		return
	}

	resp = doGetUserInfoByIds(ctx, req)
	ctx.JSON(http.StatusOK, resp)
}

func doGetUserInfoByIds(ctx *gin.Context, req *pb.GetUserInfoByIdsReq) (resp *pb.GetUserInfoByIdsResp) {
	users, err := getUserInfoByIds(req.GetIds())
	if err != nil {
		resp.Base = pkg.InternalError(err)
	} else {
		resp.Base = pkg.Success()
		resp.Infos = users
	}
	return resp
}

func getUserInfoByIds(uids []uint64) ([]*pb.User, error) {
	users, err := domain.UserService.GetByIds(uids)
	if err != nil {
		return nil, err
	}
	vos := userMoToVo(users)
	return vos, nil
}

func userMoToVo(users []*model.User) []*pb.User {
	m := map[bool]pb.UserState{true: pb.UserState_ON, false: pb.UserState_OFF}
	vos := make([]*pb.User, len(users))
	for i, user := range users {
		state := m[domain.ConnService.GetUserState(user.ID)]
		vo := &pb.User{
			Id:       user.ID,
			Username: user.Username,
			Nickname: user.Nickname,
			State:    state,
		}
		vos[i] = vo
	}
	return vos
}
