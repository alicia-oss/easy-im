package pkg

import "errors"

const (
	ModuleNameRepo    = "domain_user_repo"
	ModuleNameService = "domain_user_service"
)

var (
	ErrUnknown       = errors.New("unknown internal error")
	ErrUserNotExist  = errors.New("user not exist")
	ErrWrongPassword = errors.New("wrong password")
)
