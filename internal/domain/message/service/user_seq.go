package service

import (
	"easy_im/internal/domain/message/model"
	"easy_im/internal/domain/message/pkg"
	"easy_im/internal/domain/message/repo"
	"easy_im/pkg/log"
	"fmt"
)

var UserService = NewUserSeqServiceImpl()

type IUserSeqService interface {
	GenSeq(userId uint64) (uint64, error)
	CreateSeqBox(userId uint64) error
}

type userSeqServiceImpl struct {
	lock        pkg.IUserSeqLock
	userSeqRepo repo.IUserSeqRepo
}

func NewUserSeqServiceImpl() IUserSeqService {
	return &userSeqServiceImpl{
		lock:        pkg.NewUserSeqLock(),
		userSeqRepo: repo.NewUserSeqRepo(),
	}
}

func (u *userSeqServiceImpl) CreateSeqBox(userId uint64) error {
	seq := &model.UserSeq{
		UserId: userId,
		MaxSeq: 100,
	}
	err := u.userSeqRepo.Add(seq)
	if err != nil {
		return pkg.ErrUnknown
	}
	return nil
}

func (u *userSeqServiceImpl) GenSeq(userId uint64) (uint64, error) {
	seq, max, err := u.userSeqRepo.RGetUserSeq(userId)
	if err != nil {
		return 0, pkg.ErrUnknown
	}
	if seq != 0 {
		return seq, nil
	}

	info := fmt.Sprintf("redis user_seq:%v not ready, will load from db", userId)
	log.Info(info, pkg.ModuleNameRepoUserSeq)

	err = u.resetSeq(userId, max)
	if err != nil {
		return 0, err
	}

	seq, _, err = u.userSeqRepo.RGetUserSeq(userId)
	if seq == 0 || err != nil {
		log.Error(fmt.Sprintf("GenSeq user_id:%v err: load error", userId), pkg.ModuleNameServiceUserSeq)
		return 0, pkg.ErrUnknown
	}
	return seq, nil

}

func (u *userSeqServiceImpl) resetSeq(userId, preMax uint64) error {
	u.lock.Lock(userId)
	defer u.lock.UnLock(userId)

	userReq, err := u.userSeqRepo.GetByUserId(userId)
	if err != nil {
		return pkg.ErrUnknown
	}
	if preMax != userReq.MaxSeq {
		return nil
	}

	curSeq := userReq.MaxSeq
	userReq.MaxSeq += 100
	err = u.userSeqRepo.Save(userReq)
	if err != nil {
		return pkg.ErrUnknown
	}
	err = u.userSeqRepo.RSetUserSeq(userId, curSeq, userReq.MaxSeq)
	if err != nil {
		return pkg.ErrUnknown
	}
	return nil
}
