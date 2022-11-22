package service

import (
	"easy_im/internal/domain/conn/adapter"
	"easy_im/internal/domain/conn/model"
	"easy_im/internal/domain/conn/pkg"
	"easy_im/internal/domain/conn/repo"
	"easy_im/pkg/log"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var ConnService = NewConnService(3 * time.Second)

type IConnService interface {
	Run()
	Stop()
	SendNotice(userId uint64, noticeType int32, data []byte) error
	SendAck(userId uint64, id string, data []byte) error
	OnlineUser(userId uint64, conn adapter.IConnection)
	OfflineUser(userId uint64) error
	GetUserState(userId uint64) bool
}

func NewConnService(retryTime time.Duration) IConnService {
	return &connServiceImpl{
		connManager:   repo.NewConnManager(),
		retryQueue:    pkg.NewMsgRetryQueue(),
		noticeRepo:    repo.NewNoticeRepo(),
		retryDuration: retryTime,
		closeChan:     make(chan struct{}),
	}
}

type connServiceImpl struct {
	connManager   repo.IConnManager
	retryQueue    *pkg.MsgRetryQueue
	noticeRepo    repo.INoticeRepo
	retryDuration time.Duration
	closeChan     chan struct{}
}

func (r *connServiceImpl) Run() {
	go r.retryQueue.Run()
	go r.startListener()
}

func (r *connServiceImpl) Stop() {
	r.retryQueue.Close()
}

func (r *connServiceImpl) SendNotice(userId uint64, noticeType int32, data []byte) error {
	// gen notice
	n := &model.Notice{
		ID:         uuid.New().String(),
		UserId:     userId,
		NoticeType: noticeType,
		Data:       data,
	}
	// save notice
	err := r.noticeRepo.RAdd(n)
	if err != nil {
		return pkg.ErrUnknown
	}
	// send notice
	return r.doSendNotice(n)

}

func (r *connServiceImpl) SendAck(userId uint64, id string, data []byte) error {
	conn, ok := r.connManager.GetConnById(userId)
	if !ok {
		return pkg.ReceiverOffline
	}
	err := conn.SendAck(data, id)
	if err != nil {
		return pkg.ErrUnknown
	}
	return nil

}

func (r *connServiceImpl) OnlineUser(userId uint64, conn adapter.IConnection) {
	c := model.NewConn(userId, conn, r.connManager)
	c.Online()
}

func (r *connServiceImpl) OfflineUser(userId uint64) error {
	c, ok := r.connManager.GetConnById(userId)
	if !ok {
		return pkg.UserOffline
	}
	c.Offline()
	return nil
}

func (r *connServiceImpl) doSendNotice(notice *model.Notice) error {
	conn, ok := r.connManager.GetConnById(notice.UserId)
	if !ok {
		return pkg.ReceiverOffline
	}
	err := conn.SendNotice(notice.Data, notice.NoticeType, notice.ID)
	if err != nil {
		log.Error(fmt.Sprintf("do send notice err:%v, notice_id:%v", err, notice.ID), "domain_conn_retry_manager")
	}
	r.addRetry(notice.ID)
	return nil
}

func (r *connServiceImpl) startListener() {
	for {
		noticeId, ok := <-r.retryQueue.Listen()
		if !ok {
			log.Info("listener stop.....", "domain_conn_retry_manager")
			return
		}
		t, err := r.noticeRepo.RGet(noticeId)
		if err != nil || t == nil {
			continue
		}

		_ = r.doSendNotice(t)
	}
}

func (r *connServiceImpl) addRetry(noticeId string) {
	r.retryQueue.SubmitJob(&pkg.Job{
		Info:     noticeId,
		Duration: r.retryDuration,
	})
}

func (r *connServiceImpl) GetUserState(userId uint64) bool {
	_, b := r.connManager.GetConnById(userId)
	return b
}
