package repo

import (
	"easy_im/internal/domain/im/model"
	"easy_im/internal/domain/im/pkg"
	"easy_im/internal/domain/im/repo/lua"
	"easy_im/pkg/db"
	"easy_im/pkg/log"
	"easy_im/pkg/redis"
	"fmt"
	"gorm.io/gorm"
)

type IUserSeqRepo interface {
	Add(seq *model.UserSeq) error
	Get(id uint) (*model.UserSeq, error)
	Save(seq *model.UserSeq) error
	Delete(id uint) error
	GetByUserId(userId uint) (*model.UserSeq, error)

	RSetUserSeq(userId uint, curSeq uint, maxSeq uint) error
	RGetUserSeq(userId uint) (uint, error)
}

func NewUserRepo() IUserSeqRepo {
	return &userSeqRepoImpl{}
}

type userSeqRepoImpl struct{}

func (i *userSeqRepoImpl) RSetUserSeq(userId uint, curSeq uint, maxSeq uint) error {
	key := model.BuildUserSeqKey(userId)
	value := model.BuildSeqValue(curSeq, maxSeq)
	_, err := redis.Client.Set(key, value, model.GetSeqTTL()).Result()
	if err != nil {
		log.Error(fmt.Sprintf("RSetUserSeq error:%v", err), pkg.ModuleNameRepoUserSeq)
	}
	return err
}

func (i *userSeqRepoImpl) RGetUserSeq(userId uint) (uint, error) {
	return lua.GetSeq(model.BuildUserSeqKey(userId))
}

func (*userSeqRepoImpl) Add(seq *model.UserSeq) error {
	err := db.DB.Create(seq).Error
	if err != nil {
		log.Error(fmt.Sprintf("add error :%v", err), pkg.ModuleNameRepoUserSeq)
		return err
	}
	return nil
}

// Get 获取用户信息
func (*userSeqRepoImpl) Get(id uint) (*model.UserSeq, error) {
	var seq = model.UserSeq{
		Model: gorm.Model{ID: id},
	}
	err := db.DB.First(&seq).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error(fmt.Sprintf("get error :%v", err), pkg.ModuleNameRepoUserSeq)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &seq, nil
}

// Save 保存
func (*userSeqRepoImpl) Save(seq *model.UserSeq) error {
	err := db.DB.Save(seq).Error
	if err != nil {
		log.Error(fmt.Sprintf("save error :%v", err), pkg.ModuleNameRepoUserSeq)
		return err
	}
	return nil
}

func (*userSeqRepoImpl) Delete(id uint) error {
	if err := db.DB.Delete(&model.UserSeq{}, id).Error; err != nil {
		log.Error(fmt.Sprintf("Delete error :%v", err), pkg.ModuleNameRepoUserSeq)
		return err
	}
	return nil
}

// GetByUserId 根据用户名获取用户信息
func (*userSeqRepoImpl) GetByUserId(userId uint) (*model.UserSeq, error) {
	var user model.UserSeq
	err := db.DB.First(&user, "user_id = ?", userId).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Error(fmt.Sprintf("get error :%v", err), pkg.ModuleNameRepoUserSeq)
		return nil, err
	}
	return &user, nil
}
