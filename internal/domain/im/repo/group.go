package repo

import (
	"easy_im/internal/domain/im/model"
	"easy_im/internal/domain/im/pkg"
	"easy_im/pkg/db"
	"easy_im/pkg/log"
	"fmt"
	"gorm.io/gorm"
)

type IGroupRepo interface {
	Add(group *model.Group) error
	Get(id uint) (*model.Group, error)
	GetByIds(ids []uint) ([]*model.Group, error)
	Save(group *model.Group) error
	Delete(id uint) error
}

type groupRepoImpl struct{}

func (*groupRepoImpl) GetByIds(ids []uint) ([]*model.Group, error) {
	var res []*model.Group
	err := db.DB.Where(ids).Find(res).Error
	if err != nil {
		log.Error(fmt.Sprintf("GetByIds error :%v", err), pkg.ModuleNameRepoGroup)
		return nil, err
	}
	return res, nil
}

func (*groupRepoImpl) Add(group *model.Group) error {
	err := db.DB.Create(group).Error
	if err != nil {
		log.Error(fmt.Sprintf("add error :%v", err), pkg.ModuleNameRepoGroup)
		return err
	}
	return nil
}

func (*groupRepoImpl) Get(id uint) (*model.Group, error) {
	var group = model.Group{
		Model: gorm.Model{ID: id},
	}
	err := db.DB.First(&group).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error(fmt.Sprintf("get error :%v", err), pkg.ModuleNameRepoGroup)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &group, nil
}

func (*groupRepoImpl) Save(group *model.Group) error {
	err := db.DB.Save(group).Error
	if err != nil {
		log.Error(fmt.Sprintf("save error :%v", err), pkg.ModuleNameRepoGroup)
		return err
	}
	return nil
}

func (*groupRepoImpl) Delete(id uint) error {
	if err := db.DB.Delete(&model.Group{}, id).Error; err != nil {
		log.Error(fmt.Sprintf("Delete error :%v", err), pkg.ModuleNameRepoGroup)
		return err
	}
	return nil
}
