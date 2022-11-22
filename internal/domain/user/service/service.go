package service

import (
	"easy_im/internal/domain/user/model"
	"easy_im/internal/domain/user/pkg"
	"easy_im/internal/domain/user/repo"
)

var UserService IUserService

func init() {
	UserService = &userServiceImpl{userRepo: repo.NewUserRepo()}
}

type IUserService interface {
	Auth(username, password string) (error, *model.User)
	Register(username, password, nickname string) (*model.User, error)
	Update(user *model.User) error
	Logout(userId uint64) error
	GetById(userId uint64) (*model.User, error)
	GetByIds(ids []uint64) ([]*model.User, error)
	Search(key string) ([]*model.User, error)
}

type userServiceImpl struct {
	userRepo repo.IUserRepo
}

func (u *userServiceImpl) GetById(userId uint64) (*model.User, error) {
	user, err := u.userRepo.Get(userId)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	return user, nil
}

func (u *userServiceImpl) GetByIds(ids []uint64) ([]*model.User, error) {
	user, err := u.userRepo.GetByIds(ids)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	return user, nil
}

func (u *userServiceImpl) Search(key string) ([]*model.User, error) {
	user, err := u.userRepo.Search(key)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	return user, nil
}

func (u *userServiceImpl) Auth(username, password string) (err error, m *model.User) {
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return pkg.ErrUnknown, nil
	}
	if user == nil {
		return pkg.ErrUserNotExist, nil
	}
	if user.Password != password {
		return pkg.ErrWrongPassword, nil
	}
	return nil, user
}

func (u *userServiceImpl) Register(username, password, nickname string) (*model.User, error) {
	user := &model.User{
		Username: username,
		Password: password,
		Nickname: nickname,
	}
	t, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	if t != nil {
		return nil, pkg.ErrUsernameUsed
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

func (u *userServiceImpl) Logout(userId uint64) error {
	if err := u.userRepo.Delete(userId); err != nil {
		return pkg.ErrUnknown
	}
	return nil
}
