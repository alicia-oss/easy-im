package user_imp

import (
	"easy_im/service/user_manager/int"
	"sync"
)

func NewUserManager() user_int.IUserManager {
	return &userManager{userMap: &sync.Map{}}
}

type userManager struct {
	userMap *sync.Map
}

func (u *userManager) GetUserById(userId uint32) (user_int.IUser, bool) {
	load, ok := u.userMap.Load(userId)
	if !ok {
		return nil, ok
	}
	return load.(user_int.IUser), ok
}

func (u *userManager) OfflineAllUser() {

}

func (u *userManager) GetAllUser() []user_int.IUser {
	var res []user_int.IUser
	u.userMap.Range(func(key, value any) bool {
		user := value.(user_int.IUser)
		res = append(res, user)
		return true
	})
	return res
}

func (u *userManager) GetUserNum() int {
	res := 0
	u.userMap.Range(func(key, value any) bool {
		res++
		return true
	})
	return res
}

func (u *userManager) AddUser(user user_int.IUser) {
	u.userMap.Store(user.GetUserId(), user)
}

func (u *userManager) RemoveUser(userId uint32) {
	u.userMap.Delete(userId)
}
