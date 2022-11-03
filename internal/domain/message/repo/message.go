package repo

import (
	"easy_im/internal/domain/message/model"
	"easy_im/internal/domain/message/pkg"
	"easy_im/pkg/db"
	"easy_im/pkg/log"
	"easy_im/pkg/redis"
	"fmt"
	redis2 "github.com/go-redis/redis"
	"gorm.io/gorm"
	"strconv"
)

type IMessageRepo interface {
	Add(entity *model.Message) error
	Get(id uint) (*model.Message, error)
	GetByIds(ids []uint) ([]*model.Message, error)
	Save(entity *model.Message) error
	Delete(entity *model.Message) error

	RAddUserInbox(userId uint, msg *model.Message) error
	RRangeGetUserInbox(userId uint, senderId uint, min, max uint) ([]uint, error)
	RGetUserInbox(userId uint) ([]uint, error)
}

type messageRepoImpl struct{}

func (i *messageRepoImpl) RAddUserInbox(userId uint, msg *model.Message) error {
	key := model.BuildInboxKey(userId)
	core := model.BuildInboxCore(msg.SenderId, msg.Seq)
	err := redis.Client.ZAdd(key, redis2.Z{
		Score:  core,
		Member: msg.ID,
	}).Err()
	if err != nil {
		log.Error(fmt.Sprintf("RAddUserInbox error:%v", err), pkg.ModuleNameRepoMessage)
		return err
	}
	return nil
}

func (i *messageRepoImpl) RGetUserInbox(userId uint) ([]uint, error) {
	key := model.BuildInboxKey(userId)
	result, err := redis.Client.ZRange(key, 0, -1).Result()
	if err != nil {
		log.Error(fmt.Sprintf("RGetUserInbox error:%v", err), pkg.ModuleNameRepoMessage)
		return nil, err
	}
	res, err := stringToUint(result)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (i *messageRepoImpl) RRangeGetUserInbox(userId uint, senderId uint, min, max uint) ([]uint, error) {
	key := model.BuildInboxKey(userId)
	minCore := model.BuildInboxCore(senderId, min)
	maxCore := model.BuildInboxCore(senderId, max)
	result, err := redis.Client.ZRange(key, int64(minCore), int64(maxCore)).Result()
	if err != nil {
		log.Error(fmt.Sprintf("RRangeGetUserInbox error:%v", err), pkg.ModuleNameRepoMessage)
		return nil, err
	}
	res, err := stringToUint(result)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (*messageRepoImpl) Add(entity *model.Message) error {
	err := db.DB.Create(entity).Error
	if err != nil {
		log.Error(fmt.Sprintf("add error :%v", err), pkg.ModuleNameRepoMessage)
		return err
	}
	return nil
}

func (*messageRepoImpl) Get(id uint) (*model.Message, error) {
	var entity = model.Message{
		Model: gorm.Model{ID: id},
	}
	err := db.DB.First(&entity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error(fmt.Sprintf("get error :%v", err), pkg.ModuleNameRepoMessage)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &entity, nil
}

func (*messageRepoImpl) GetByIds(ids []uint) ([]*model.Message, error) {
	var res []*model.Message
	err := db.DB.Where(ids).Find(res).Error
	if err != nil {
		log.Error(fmt.Sprintf("GetByIds error :%v", err), pkg.ModuleNameRepoMessage)
		return nil, err
	}
	return res, nil
}

func (*messageRepoImpl) Save(entity *model.Message) error {
	err := db.DB.Save(entity).Error
	if err != nil {
		log.Error(fmt.Sprintf("save error :%v", err), pkg.ModuleNameRepoMessage)
		return err
	}
	return nil
}

func (*messageRepoImpl) Delete(entity *model.Message) error {
	if err := db.DB.Delete(entity).Error; err != nil {
		log.Error(fmt.Sprintf("Delete error :%v", err), pkg.ModuleNameRepoMessage)
		return err
	}
	return nil
}

func stringToUint(result []string) ([]uint, error) {
	res := make([]uint, len(result))
	for i, s := range result {
		t, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			log.Error(fmt.Sprintf("RRangeGetUserInbox ParseUint error:%v", err), pkg.ModuleNameRepoMessage)
			return nil, err
		}
		res[i] = uint(t)
	}
	return res, nil
}
