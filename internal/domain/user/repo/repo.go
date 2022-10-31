package repo

import (
	"easy_im/internal/domain/user/model"
	"easy_im/internal/domain/user/pkg"
	"easy_im/pkg/db"
	"easy_im/pkg/log"
	"fmt"
	"gorm.io/gorm"
)

var Repo = NewUserRepo()

type IUserRepo interface {
	Add(user *model.User) error
	Get(userId uint) (*model.User, error)
	Save(user *model.User) error
	Delete(userId uint) error
	GetByUsername(username string) (*model.User, error)
	GetByIds(userIds []int64) ([]*model.User, error)
	Search(key string) ([]*model.User, error)
}

type userRepoImpl struct {
}

func NewUserRepo() IUserRepo {
	return &userRepoImpl{}
}

func (*userRepoImpl) Add(user *model.User) error {
	err := db.DB.Create(user).Error
	if err != nil {
		log.Error(fmt.Sprintf("add error :%v", err), pkg.ModuleNameRepo)
		return err
	}
	return nil
}

// Get 获取用户信息
func (*userRepoImpl) Get(userId uint) (*model.User, error) {
	var user = model.User{
		Model: gorm.Model{ID: userId},
	}
	err := db.DB.First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error(fmt.Sprintf("get error :%v", err), pkg.ModuleNameRepo)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, nil
}

// Save 保存
func (*userRepoImpl) Save(user *model.User) error {
	err := db.DB.Save(user).Error
	if err != nil {
		log.Error(fmt.Sprintf("save error :%v", err), pkg.ModuleNameRepo)
		return err
	}
	return nil
}

func (*userRepoImpl) Delete(userId uint) error {
	if err := db.DB.Delete(&model.User{}, userId).Error; err != nil {
		log.Error(fmt.Sprintf("Delete error :%v", err), pkg.ModuleNameRepo)
		return err
	}
	return nil
}

// GetByUsername 根据用户名获取用户信息
func (*userRepoImpl) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := db.DB.First(&user, "username = ?", username).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Error(fmt.Sprintf("save error :%v", err), pkg.ModuleNameRepo)
		return nil, err
	}
	return &user, nil
}

// GetByIds 获取用户信息
func (*userRepoImpl) GetByIds(userIds []int64) ([]*model.User, error) {
	var users []*model.User
	err := db.DB.Find(&users, "id in (?)", userIds).Error
	if err != nil {
		log.Error(fmt.Sprintf("GetByIds error :%v", err), pkg.ModuleNameRepo)
		return nil, err
	}
	return users, nil
}

// Search 查询用户,这里简单实现，生产环境建议使用ES
func (*userRepoImpl) Search(key string) ([]*model.User, error) {
	var users []*model.User
	key = "%" + key + "%"
	err := db.DB.Where("username like ? or nickname like ?", key, key).Find(&users).Error
	if err != nil {
		log.Error(fmt.Sprintf("Search error :%v", err), pkg.ModuleNameRepo)
		return nil, err
	}
	return users, nil
}
