package pkg

import "errors"

var (
	ErrUnknown      = errors.New("unknown internal error")
	ReceiverOffline = errors.New("receiver not online")
	UserOffline     = errors.New("user not online")
)
