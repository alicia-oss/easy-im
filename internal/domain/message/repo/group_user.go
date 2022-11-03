package repo

import (
	"easy_im/internal/domain/message/model"
	"easy_im/internal/domain/message/pkg"
	"easy_im/pkg/db"
	"easy_im/pkg/log"
	"fmt"
	"gorm.io/gorm"
)

type IGroupUserRepo interface {
	Add(gu *model.GroupUser) error
	Get(id uint64) (*model.GroupUser, error)
	Save(gu *model.GroupUser) error
	Delete(gu *model.GroupUser) error
	GetByGroupId(id uint64) ([]*model.GroupUser, error)
	GetByUserId(id uint64) ([]*model.GroupUser, error)
}

func NewGroupUserRepo() IGroupUserRepo {
	return &groupUserRepoImpl{}
}

type groupUserRepoImpl struct{}

func (*groupUserRepoImpl) Add(gu *model.GroupUser) error {
	err := db.DB.Create(gu).Error
	if err != nil {
		log.Error(fmt.Sprintf("add error :%v", err), pkg.ModuleNameRepoGroupUser)
		return err
	}
	return nil
}

func (*groupUserRepoImpl) Get(id uint64) (*model.GroupUser, error) {
	var seq = model.GroupUser{
		ID: id,
	}
	err := db.DB.First(&seq).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error(fmt.Sprintf("get error :%v", err), pkg.ModuleNameRepoGroupUser)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &seq, nil
}

func (*groupUserRepoImpl) Save(gu *model.GroupUser) error {
	err := db.DB.Save(gu).Error
	if err != nil {
		log.Error(fmt.Sprintf("save error :%v", err), pkg.ModuleNameRepoGroupUser)
		return err
	}
	return nil
}

func (*groupUserRepoImpl) Delete(gu *model.GroupUser) error {
	if err := db.DB.Where(gu).Delete(&model.GroupUser{}).Error; err != nil {
		log.Error(fmt.Sprintf("Delete error :%v", err), pkg.ModuleNameRepoGroupUser)
		return err
	}
	return nil
}

func (*groupUserRepoImpl) GetByGroupId(id uint64) ([]*model.GroupUser, error) {
	var res []*model.GroupUser
	err := db.DB.Where(&model.GroupUser{GroupId: id}).Find(res).Error
	if err != nil {
		log.Error(fmt.Sprintf("GetByGroupId error :%v", err), pkg.ModuleNameRepoGroupUser)
		return nil, err
	}
	return res, nil
}

func (*groupUserRepoImpl) GetByUserId(id uint64) ([]*model.GroupUser, error) {
	var res []*model.GroupUser
	err := db.DB.Where(&model.GroupUser{UserId: id}).Find(res).Error
	if err != nil {
		log.Error(fmt.Sprintf("GetByUserId error :%v", err), pkg.ModuleNameRepoGroupUser)
		return nil, err
	}
	return res, nil
}
