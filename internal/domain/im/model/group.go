package model

import (
	"easy_im/pkg/db"
	"easy_im/pkg/log"
	"fmt"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	GroupName string `gorm:"comment:群名称;type:varchar(256);not null;"`
	OwnerId   uint   `gorm:"comment:群主;not null;"`
	CreatorId uint   `gorm:"comment:群的创造者Id;not null;"`
	Extra     string `gorm:"comment:拓展信息;type:varchar(256);"`
}

func (*Group) TableName() string { return "group" }

type GroupUser struct {
	gorm.Model
	GroupId uint `gorm:""`
	UserId  uint `gorm:""`
}

func (*GroupUser) TableName() string { return "group_user" }

func init() {
	err := db.DB.AutoMigrate(&Group{}, &GroupUser{})
	if err != nil {
		log.Error(fmt.Sprintf("group, group_user init err:%v", err), "domain_user_model")
		panic(err.Error())
	}
}
