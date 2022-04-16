package router

import (
	"net/http"
)

var defaultUserRouter *UserRouter

// GetUserRouter 获取用户中心根路由
func GetUserRouter() *UserRouter {
	return defaultUserRouter
}

func init() {
	defaultUserRouter = &UserRouter{}
}

// UserRouter 用户中心根路由
type UserRouter struct {
	path    string
	handler http.HandlerFunc
}

// GetPath 获取路径
func (u *UserRouter) GetPath() string {
	return u.path
}

// GetHandler 获取路由处理器
func (u *UserRouter) GetHandler() http.HandlerFunc {
	return u.handler
}

// CreateUserRouter 创建用户的router
func (u *UserRouter) CreateUserRouter() Router {
	u.path = "/user/create"
	u.handler = createUserHandler
	return u
}

func createUserHandler(resp http.ResponseWriter, req *http.Request) {
	println("/user/create is in")
	// 这里可以做一些数据合法性校验，如os注入攻击
	// 转controller处理
	// 处理返回的结果
}

// Router 路由抽象类，解耦路由
type Router interface {
	GetPath() string
	GetHandler() http.HandlerFunc
}
