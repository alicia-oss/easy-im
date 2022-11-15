package repo

import (
	"easy_im/internal/domain/conn/model"
	"sync"
)

var ConnManager = NewConnManager()

type IConnManager interface {
	GetConnById(userId uint64) (model.IConn, bool)
	OfflineAllConn()
	GetAllConn() []model.IConn
	GetConnNum() int
	AddConn(user model.IConn)
	RemoveConn(userId uint64)
}

func NewConnManager() IConnManager {
	return &connManagerImpl{userMap: &sync.Map{}}
}

type connManagerImpl struct {
	userMap *sync.Map
}

func (u *connManagerImpl) GetConnById(userId uint64) (model.IConn, bool) {
	load, ok := u.userMap.Load(userId)
	if !ok {
		return nil, ok
	}
	return load.(model.IConn), ok
}

func (u *connManagerImpl) OfflineAllConn() {
	//TODO implement me
	panic("implement me")
}

func (u *connManagerImpl) GetConnNum() int {
	res := 0
	u.userMap.Range(func(key, value any) bool {
		res++
		return true
	})
	return res
}

func (u *connManagerImpl) AddConn(conn model.IConn) {
	u.userMap.Store(conn.GetUserId(), conn)

}

func (u *connManagerImpl) RemoveConn(userId uint64) {
	u.userMap.Delete(userId)
}

func (u *connManagerImpl) GetAllConn() []model.IConn {
	var res []model.IConn
	u.userMap.Range(func(key, value any) bool {
		user := value.(model.IConn)
		res = append(res, user)
		return true
	})
	return res
}
