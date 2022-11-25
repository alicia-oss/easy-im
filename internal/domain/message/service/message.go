package service

import (
	"easy_im/internal/domain/message/model"
	"easy_im/internal/domain/message/pkg"
	"easy_im/internal/domain/message/repo"
	"easy_im/pkg/log"
	"time"
)

func NewMessageService() IMessageService {
	return &messageServiceImpl{
		msgRepo: repo.NewMessageRepo(),
	}
}

type IMessageService interface {
	// SaveMessage  处理用户要发送的消息
	SaveMessage(message *model.Message) error
	// SyncInbox 同步收件箱
	SyncInbox(userId, senderId uint64, min uint64, max uint64) ([]*model.Message, error)
	// GetMessageIdsByUserId 获取用户的全部消息
	GetMessageIdsByUserId(userId uint64) ([]uint64, error)
	// GetMessageByIds 获取消息
	GetMessageByIds(ids []uint64) ([]*model.Message, error)

	UpdateMsgStateToSeen(uid uint64, mid uint64, t time.Time) error
	UpdateMsgStateToDelivered(mid uint64, uid uint64, t time.Time) error
}

type messageServiceImpl struct {
	msgRepo repo.IMessageRepo
}

func (m *messageServiceImpl) UpdateMsgStateToSeen(mid uint64, uid uint64, t time.Time) error {
	message, err := m.msgRepo.Get(mid)
	if err != nil {
		return pkg.ErrUnknown
	}
	if message == nil {
		return pkg.MsgNotExist
	}
	if message.ReceiverId != uid {
		return pkg.InvalidUserId
	}
	message.State = pkg.MessageState_SEEN
	err = m.msgRepo.Save(message)
	if err != nil {
		return pkg.ErrUnknown
	}
	return nil
}

func (m *messageServiceImpl) UpdateMsgStateToDelivered(mid uint64, uid uint64, t time.Time) error {
	message, err := m.msgRepo.Get(mid)
	if err != nil {
		return pkg.ErrUnknown
	}
	if message == nil {
		return pkg.MsgNotExist
	}
	if message.ReceiverId != uid {
		return pkg.InvalidUserId
	}
	message.State = pkg.MessageState_DELIVERED
	err = m.msgRepo.Save(message)
	if err != nil {
		return pkg.ErrUnknown
	}
	return nil
}

func (m *messageServiceImpl) SaveMessage(message *model.Message) error {
	err := m.msgRepo.Save(message)
	if err != nil {
		return pkg.ErrUnknown
	}
	err = m.msgRepo.RAddUserInbox(message.ReceiverId, message)
	if err != nil {
		return pkg.ErrUnknown
	}
	// 写入kafka
	return nil
}

func (m *messageServiceImpl) SyncInbox(userId, senderId uint64, min uint64, max uint64) ([]*model.Message, error) {
	ids, err := m.msgRepo.RRangeGetUserInbox(userId, senderId, min, max)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	return m.GetMessageByIds(ids)
}

func (m *messageServiceImpl) GetMessageIdsByUserId(userId uint64) ([]uint64, error) {
	ids, err := m.msgRepo.RGetUserInbox(userId)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	return ids, nil
}

func (m *messageServiceImpl) GetMessageByIds(ids []uint64) ([]*model.Message, error) {
	res, err := m.msgRepo.GetByIds(ids)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	if len(res) != len(ids) {
		log.Error("SyncInbox error: len(res) != len(ids)", pkg.ModuleNameServiceMessage)
	}
	return res, nil
}
