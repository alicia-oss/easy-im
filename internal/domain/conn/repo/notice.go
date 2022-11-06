package repo

import (
	"easy_im/internal/domain/conn/model"
	"easy_im/pkg/redis"
	"time"
)

func NewNoticeRepo() INoticeRepo {
	return &noticeRepoImpl{}
}

type INoticeRepo interface {
	RAdd(notice *model.Notice) error
	RGet(noticeId string) (*model.Notice, error)
}

type noticeRepoImpl struct{}

func (n *noticeRepoImpl) RAdd(notice *model.Notice) error {
	key := model.BuildNoticeKey(notice.ID)
	_, err := redis.Client.Set(key, notice.Data, 24*time.Hour).Result()
	return err
}

func (n *noticeRepoImpl) RGet(noticeId string) (*model.Notice, error) {
	key := model.BuildNoticeKey(noticeId)
	res, err := redis.Client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	return &model.Notice{
		ID:   noticeId,
		Data: []byte(res),
	}, nil
}
