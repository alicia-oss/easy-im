package user_int

type IUserManager interface {
	GetUserById(userId uint) (IUser, bool)
	OfflineAllUser()
	GetAllUser() []IUser
	GetUserNum() int
	AddUser(user IUser)
	RemoveUser(userId uint)
}
