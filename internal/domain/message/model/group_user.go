package model

import (
	"gorm.io/gorm"
	"time"
)

type GroupUser struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	GroupId   uint64         `gorm:""`
	UserId    uint64         `gorm:""`
}

func (*GroupUser) TableName() string { return "group_user" }
