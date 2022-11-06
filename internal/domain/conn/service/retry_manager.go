package service

import (
	"easy_im/internal/domain/conn/pkg"
	"easy_im/internal/domain/conn/repo"
	"time"
)

type IRetryManager interface {
	Run()
	Stop()
	AddNotice(noticeId string)
}

type retryManagerImpl struct {
	retryQueue *pkg.MsgRetryQueue
	noticeRepo repo.INoticeRepo
	retryDuration time.Duration
	closeChan chan struct{}
}




