package router

import (
	"net/http"
)

var defaultUserRouter *UserRouter

func GetUserRouter() *UserRouter {
	return defaultUserRouter
}

func init() {
	defaultUserRouter = &UserRouter{}
}

type UserRouter struct {
	path    string
	handler http.HandlerFunc
}

func (u *UserRouter) GetPath() string {
	return u.path
}

func (u *UserRouter) GetHandler() http.HandlerFunc {
	return u.handler
}

func (u *UserRouter) CreateUserRout() Router {
	u.path = "/user/create"
	u.handler = createUserHandler
	return u
}

func createUserHandler(resp http.ResponseWriter, req *http.Request) {
	// 这里可以做一些数据校验
	// 转controller处理
	// 处理返回的结果
}

type Router interface {
	GetPath() string
	GetHandler() http.HandlerFunc
}
