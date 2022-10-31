package service

import (
	"easy_im/internal/domain/user/model"
	"easy_im/internal/domain/user/pkg"
	"easy_im/internal/domain/user/repo"
)

type IUserService interface {
	Online(username, password string) error
	Offline(userId uint) error
	Register(username, password, nickname string) (*model.User, error)
	Update(user *model.User) error
	Logout(userId uint) error
}

type userServiceImpl struct {
	userRepo repo.IUserRepo
}

func (u *userServiceImpl) Online(username, password string) (err error) {
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return pkg.ErrUnknown
	}
	if user != nil {
		return pkg.ErrUserNotExist
	}
	if user.Password != password {
		return pkg.ErrWrongPassword
	}
	return nil
}

func (u *userServiceImpl) Offline(userId uint) error {
	return nil
}

func (u *userServiceImpl) Register(username, password, nickname string) (*model.User, error) {
	user := &model.User{
		Username: username,
		Password: password,
		Nickname: nickname,
	}
	if err := u.userRepo.Add(user); err != nil {
		return nil, pkg.ErrUnknown
	}
	return user, nil
}

func (u *userServiceImpl) Update(user *model.User) error {
	if err := u.userRepo.Save(user); err != nil {
		return pkg.ErrUnknown
	}
	return nil
}

func (u *userServiceImpl) Logout(userId uint) error {
	if err := u.userRepo.Delete(userId); err != nil {
		return pkg.ErrUnknown
	}
	return nil
}
