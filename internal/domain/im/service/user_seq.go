package service

import (
	"easy_im/internal/domain/im/pkg"
	"easy_im/internal/domain/im/repo"
	"easy_im/pkg/log"
	"fmt"
)

type IUserSeqService interface {
	GenSeq(userId uint) uint64
}

type userSeqServiceImpl struct {
	userSeqRepo repo.IUserSeqRepo
}

func (u *userSeqServiceImpl) GenSeq(userId uint) (uint64, error) {
	seq, err := u.userSeqRepo.RGetUserSeq(userId)
	if err != nil {
		return 0, pkg.ErrUnknown
	}
	if seq != 0 {
		return seq, nil
	}

	info := fmt.Sprintf("redis user_seq:%v not ready, will load from db", userId)
	log.Info(info, pkg.ModuleNameRepoUserSeq)

	userReq, err := u.userSeqRepo.GetByUserId(userId)
	if err != nil {
		return 0, pkg.ErrUnknown
	}
	curSeq := userReq.MaxSeq
	userReq.MaxSeq += 100
	err = u.userSeqRepo.Save(userReq)
	if err != nil {
		return 0, pkg.ErrUnknown
	}
	err = u.userSeqRepo.RSetUserSeq(userId, curSeq, userReq.MaxSeq)
	if err != nil {
		return 0, pkg.ErrUnknown
	}

	seq, err = u.userSeqRepo.RGetUserSeq(userId)
	if seq == 0 || err != nil {
		log.Error(fmt.Sprintf("GenSeq user_id:%v err: load error", userId), pkg.ModuleNameServiceUserSeq)
		return 0, pkg.ErrUnknown
	}
	return seq, nil

}
