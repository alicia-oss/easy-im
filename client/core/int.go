package core

type IClient interface {
	Online(userName string)
	Offline()
	ListUsers()
	SendMessage(content string, receiverId uint32)
	GetUserId() (uint32, bool)
	GetUserName() (string, bool)
	SetUserName(name string)
	SetUserId(uId uint32)
}
