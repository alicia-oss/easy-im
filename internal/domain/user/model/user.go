package model

import (
	"easy_im/pkg/db"
	"easy_im/pkg/log"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"uniqueIndex;comment:用户名;type:varchar(256);not null;"`
	Password  string         `gorm:"comment:密码;type:varchar(256);not null;"`
	Nickname  string         `gorm:"comment:昵称;"`
}

func (*User) TableName() string { return "user" }

func init() {
	err := db.DB.AutoMigrate(&User{})
	if err != nil {
		log.Error(fmt.Sprintf("user init err:%v", err), "domain_user_model")
		panic(err.Error())
	}
}
