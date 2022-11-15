package repo

import (
	"easy_im/internal/domain/conn/model"
	"easy_im/pkg/log"
	"easy_im/pkg/redis"
	"encoding/json"
	"fmt"
	"time"
)

func NewNoticeRepo() INoticeRepo {
	return &noticeRepoImpl{}
}

type INoticeRepo interface {
	RAdd(notice *model.Notice) error
	RGet(noticeId string) (*model.Notice, error)
	RDelete(noticeId string) error
}

type noticeRepoImpl struct {
	ackRecord map[string]bool
}

func (n *noticeRepoImpl) RAdd(notice *model.Notice) error {
	key := model.BuildNoticeKey(notice.ID)
	bytes, _ := json.Marshal(notice)
	err := redis.Client.Set(key, string(bytes), 24*time.Hour).Err()
	if err != nil {
		log.Error(fmt.Sprintf("RAdd err:%v", err), "domain_conn_repo")
		return err
	}
	return nil
}

func (n *noticeRepoImpl) RGet(noticeId string) (*model.Notice, error) {
	key := model.BuildNoticeKey(noticeId)
	bytes, err := redis.Client.Get(key).Result()
	if err != nil {
		log.Error(fmt.Sprintf("RGet err:%v", err), "domain_conn_repo")
		return nil, err
	}
	res := &model.Notice{}
	err = json.Unmarshal([]byte(bytes), res)
	if err != nil {
		log.Error(fmt.Sprintf("RGet err:%v", err), "domain_conn_repo")
		return nil, err
	}
	return res, nil
}

func (n *noticeRepoImpl) RDelete(noticeId string) error {
	key := model.BuildNoticeKey(noticeId)
	err := redis.Client.Del(key).Err()
	if err != nil {
		log.Error(fmt.Sprintf("RDelete err:%v", err), "domain_conn_repo")
		return err
	}
	return nil
}
