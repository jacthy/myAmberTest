// 负责业务处理，并将调用dto层功能实现持久化
package controller

import (
	"errors"
	"github.com/liaojuntao/infrastruct"
	"github.com/liaojuntao/infrastruct/repo"
)

// UserCtl 用户控制器，业务处理器
type UserCtl struct{}

// NewUserController 创建user controller
func NewUserController() *UserCtl {
	return &UserCtl{}
}

// CreateUser 创建用户业务逻辑
func (u *UserCtl) CreateUser(user *infrastruct.User) error {
	isNotExist, err := repo.GetUserRepo().NotExistByName(user.UserName)
	if err!=nil {
		return err
	}
	if !isNotExist {
		return errors.New("该用户名已存在")
	}
	return repo.GetUserRepo().Create(user)
}

func (u *UserCtl) UpdateUser() {

}

func (u *UserCtl) DeleteUser() {

}

func (u *UserCtl) GetUserById() {

}
