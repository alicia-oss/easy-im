package pkg

import "errors"

const (
	ModuleNameRepoUserSeq    = "domain_im_repo_user_seq"
	ModuleNameServiceUserSeq = "domain_im_service_user_seq"
	ModuleNameRepoGroup      = "domain_im_repo_group"
	ModuleNameRepoGroupUser  = "domain_im_repo_group_user"
	ModuleNameServiceGroup   = "domain_im_service_group"
	ModuleNameRepoMessage    = "domain_im_repo_message"
	ModuleNameServiceMessage = "domain_im_service_message"
)

const (
	MessageState_SENT = iota
	MessageState_DELIVERED
	MessageState_SEEN
)

const (
	MessageType_TXT = iota
	MessageType_IMG
)

const (
	SenderType_USER = iota
	SenderType_SYSTEM
)

const (
	ReceiverType_USER = iota
	ReceiverType_GROUP
)

var (
	ErrUnknown = errors.New("unknown internal error")
)
