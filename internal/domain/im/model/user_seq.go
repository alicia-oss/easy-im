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
	gorm.Model
	UserId uint `gorm:"not null;uniqueIndex;"`
	MaxSeq uint `gorm:"not null;"`
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

func BuildSeqValue(curSeq, maxSeq uint) string {
	return fmt.Sprintf("%v:%v", curSeq, maxSeq)
}

func BuildUserSeqKey(userId uint) string {
	return fmt.Sprintf("message:seq_user:%v", userId)
}

func BuildGroupSeqKey(groupId uint) string {
	return fmt.Sprintf("message:seq_group:%v", groupId)
}
