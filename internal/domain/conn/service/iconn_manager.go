package service

type IConnManager interface {
	GetConnById(userId uint64) (IConn, bool)
	OfflineAllConn()
	GetAllConn() []IConn
	GetConnNum() int
	AddConn(user IConn)
	RemoveConn(userId uint64)
}
