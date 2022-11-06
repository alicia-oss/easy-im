package service

import (
	"sync"
)

var ConnManager = NewConnManager()


func NewConnManager() IConnManager {
	return &connManagerImpl{userMap: &sync.Map{}}
}

type connManagerImpl struct {
	userMap *sync.Map
}

func (u *connManagerImpl) GetConnById(userId uint64) (IConn, bool) {
	load, ok := u.userMap.Load(userId)
	if !ok {
		return nil, ok
	}
	return load.(IConn), ok
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

func (u *connManagerImpl) AddConn(conn IConn) {
	u.userMap.Store(conn.GetUserId(), conn)

}

func (u *connManagerImpl) RemoveConn(userId uint64) {
	u.userMap.Delete(userId)
}

func (u *connManagerImpl) GetAllConn() []IConn {
	var res []IConn
	u.userMap.Range(func(key, value any) bool {
		user := value.(IConn)
		res = append(res, user)
		return true
	})
	return res
}


