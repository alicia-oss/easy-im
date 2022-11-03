package model

import (
	"easy_im/pkg/db"
	"easy_im/pkg/log"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	GroupName string         `gorm:"comment:群名称;type:varchar(256);not null;"`
	OwnerId   uint64         `gorm:"comment:群主;not null;"`
	CreatorId uint64         `gorm:"comment:群的创造者Id;not null;"`
	Extra     string         `gorm:"comment:拓展信息;type:varchar(256);"`
}

func (*Group) TableName() string { return "group" }

func init() {
	err := db.DB.AutoMigrate(&Group{}, &GroupUser{})
	if err != nil {
		log.Error(fmt.Sprintf("group, group_user init err:%v", err), "domain_user_model")
		panic(err.Error())
	}
}
