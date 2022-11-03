package user_int

type IUser interface {
	GetUserId() uint32
	GetUserName() string
	GetUserIP() string
	Offline()
	Online()
	SendMessage(data []byte, msgType uint32) error
}
