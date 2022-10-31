package pkg

import "errors"

const (
	ModuleNameRepoUserSeq    = "domain_im_repo_user_seq"
	ModuleNameServiceUserSeq = "domain_im_service_user_seq"
	ModuleNameRepoGroup      = "domain_im_repo_group"
	ModuleNameServiceGroup   = "domain_im_service_group"
	ModuleNameRepoMessage    = "domain_im_repo_message"
	ModuleNameServiceMessage = "domain_im_service_message"
)

var (
	ErrUnknown = errors.New("unknown internal error")
)
