package model

import (
	"easy_im/pkg/db"
	"easy_im/pkg/log"
	"fmt"
	"gorm.io/gorm"
	"time"
)

//mysql

type UserSeq struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserId    uint64         `gorm:"not null;uniqueIndex;"`
	MaxSeq    uint64         `gorm:"not null;"`
}

func (*UserSeq) TableName() string { return "user_seq" }

func init() {
	err := db.DB.AutoMigrate(&UserSeq{})
	if err != nil {
		log.Error(fmt.Sprintf("user_seq init err:%v", err), "domain_user_seq_model")
	}
}

// redis
//{
//	type:string
//	key: message:seq_user:user_id  ( message:seq_group:group_id)
//	value : cur_seq:max_seq
//}

func GetSeqTTL() time.Duration {
	return 7 * 24 * time.Hour
}

func BuildSeqValue(curSeq, maxSeq uint64) string {
	return fmt.Sprintf("%v:%v", curSeq, maxSeq)
}

func BuildUserSeqKey(userId uint64) string {
	return fmt.Sprintf("message:seq_user:%v", userId)
}

func BuildGroupSeqKey(groupId uint64) string {
	return fmt.Sprintf("message:seq_group:%v", groupId)
}
