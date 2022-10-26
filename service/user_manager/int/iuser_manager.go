package user_int

type IUserManager interface {
	GetUserById(userId uint32) (IUser, bool)
	OfflineAllUser()
	GetAllUser() []IUser
	GetUserNum() int
	AddUser(user IUser)
	RemoveUser(userId uint32)
}
