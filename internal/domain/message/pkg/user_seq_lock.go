package pkg

import "sync"

type IUserSeqLock interface {
	Lock(userId uint64)
	UnLock(userId uint64)
}

type userSeqLock struct {
	m    map[uint64]*sync.Mutex
	lock *sync.Mutex
}

func NewUserSeqLock() IUserSeqLock {
	return &userSeqLock{
		m:    make(map[uint64]*sync.Mutex, 32),
		lock: &sync.Mutex{},
	}
}

func (u *userSeqLock) Lock(userId uint64) {
	u.lock.Lock()

	uLock, ok := u.m[userId]
	if !ok {
		uLock = &sync.Mutex{}
		u.m[userId] = uLock
	}
	u.lock.Unlock()
	uLock.Lock()
	return

}

func (u *userSeqLock) UnLock(userId uint64) {
	u.lock.Lock()
	defer u.lock.Unlock()
	uLock := u.m[userId]
	uLock.Unlock()
}
