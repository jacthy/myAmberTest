// 负责业务处理，并将调用dto层功能实现持久化
package controller

import (
	"errors"
	"github.com/liaojuntao/infrastruct"
	"github.com/liaojuntao/infrastruct/repo"
)

// UserCtl 用户控制器，业务处理器
type UserCtl struct {
	repo infrastruct.UserRepo
}

// NewUserController 创建user controller 依赖注入方式注入repo服务，解耦
func NewUserController(userRepo infrastruct.UserRepo) *UserCtl {
	if userRepo == nil {
		return &UserCtl{repo: repo.GetSqliteUserRepo()}
	}
	return &UserCtl{repo: userRepo}
}

// CreateUser 创建用户业务逻辑
func (u *UserCtl) CreateUser(user *infrastruct.User) error {
	isNotExist, err := u.repo.NotExistByName(user.UserName)
	if err != nil {
		return err
	}
	if !isNotExist {
		return errors.New("该用户名已存在")
	}
	return u.repo.Create(user)
}

func (u *UserCtl) UpdateUser(user *infrastruct.User)error {
	existUser,err := u.repo.GetByUserName(user.UserName)
	if err != nil {
		return err
	}
	if existUser!=nil&& existUser.UserId != user.UserId {
		return errors.New("该用户名已存在")
	}
	return u.repo.Update(user)
}

func (u *UserCtl) DeleteUser() {

}

func (u *UserCtl) GetUserById() {

}
