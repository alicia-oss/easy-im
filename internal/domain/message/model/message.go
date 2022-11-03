package model

import (
	"easy_im/pkg/db"
	"easy_im/pkg/log"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID           uint64 `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	SenderType   int8
	SenderId     uint64
	ReceiverType int8
	ReceiverId   uint64
	Type         int8
	Content      []byte
	Seq          uint64
	State        int8
	SentTime     time.Time
	DeliverTime  time.Time
	SeenTime     time.Time
}

func (*Message) TableName() string { return "message" }

func init() {
	err := db.DB.AutoMigrate(&Message{})
	if err != nil {
		log.Error(fmt.Sprintf("message init err:%v", err), "domain_user_model")
		panic(err.Error())
	}
}

/*
	redis 存储 inbox
	zset
	key : message:inbox:user_id
	value: msgId
    core: sender_id*10^8 + seq

	add: zadd  message:inbox:10 1100000006 2
	getAll: zrangebyscore message:inbox:10 0 -1
	range seq: zrangebyscore message:inbox:10 100000001 1000000010
*/

func BuildInboxKey(userId uint64) string {
	return fmt.Sprintf("message:inbox:%v", userId)
}

func BuildInboxCore(senderId, seq uint64) float64 {
	return float64(senderId*100000000 + seq)
}
