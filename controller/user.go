// 负责业务处理，并将调用dto层功能实现持久化
package controller

// UserCtl 用户控制器，业务处理器
type UserCtl struct {}

// NewUserController 创建user controller
func NewUserController() *UserCtl {
	return &UserCtl{}
}

// CreateUser 创建用户业务逻辑
func (u *UserCtl) CreateUser()  {

}

func (u *UserCtl) UpdateUser()  {

}

func (u *UserCtl) DeleteUser()  {

}

func (u *UserCtl) GetUserById()  {

}