package service

import (
	"easy_im/internal/domain/im/model"
	"easy_im/internal/domain/im/pkg"
	"easy_im/internal/domain/im/repo"
	"easy_im/pkg/log"
)

type IMessageService interface {
	// SendMessage 处理用户要发送的消息
	SendMessage(message *model.Message) error
	// SyncInbox 同步收件箱
	SyncInbox(userId, senderId uint, min int, max int) ([]*model.Message, error)
	// GetMessageIdsByUserId 获取用户的全部消息
	GetMessageIdsByUserId(userId uint) ([]uint, error)
	// GetMessageByIds 获取消息
	GetMessageByIds(ids []uint) ([]*model.Message, error)
}

type messageServiceImpl struct {
	msgRepo repo.IMessageRepo
}

func (m *messageServiceImpl) SendMessage(message *model.Message) error {
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

func (m *messageServiceImpl) SyncInbox(userId, senderId uint, min uint, max uint) ([]*model.Message, error) {
	ids, err := m.msgRepo.RRangeGetUserInbox(userId, senderId, min, max)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	return m.GetMessageByIds(ids)
}

func (m *messageServiceImpl) GetMessageIdsByUserId(userId uint) ([]uint, error) {
	ids, err := m.msgRepo.RGetUserInbox(userId)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	return ids, nil
}

func (m *messageServiceImpl) GetMessageByIds(ids []uint) ([]*model.Message, error) {
	res, err := m.msgRepo.GetByIds(ids)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	if len(res) != len(ids) {
		log.Error("SyncInbox error: len(res) != len(ids)", pkg.ModuleNameServiceMessage)
	}
	return res, nil
}
