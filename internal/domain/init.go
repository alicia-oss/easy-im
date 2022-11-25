package domain

import (
	"easy_im/internal/domain/conn/service"
	service2 "easy_im/internal/domain/message/service"
	service3 "easy_im/internal/domain/user/service"
	"time"
)

var ConnService = service.NewConnService(3 * time.Second)

var GroupService = service2.NewGroupService()

var MessageService = service2.NewMessageService()

var UserSeqService = service2.NewUserSeqServiceImpl()

var UserService = service3.NewUserService()
