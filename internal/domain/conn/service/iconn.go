package service

type IConn interface {
	GetUserId() uint64
	GetUserIP() string
	Offline()
	Online()
	SendNotice(data []byte, nType uint32, noticeId uint64) error
	SendAck(data []byte, msgType uint32) error
}
