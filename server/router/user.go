package router

import (
	"github.com/liaojuntao/controller"
	"net/http"
)

const (
	statusOK = 2000 // 成功响应
	paramErr = 4001 // 参数校验错误
	optErr   = 4002 // 业务操作错误
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

// UpdateUserRouter 更新用户的router
func (u *UserRouter) UpdateUserRouter() Router {
	u.path = "/user/update"
	u.handler = updateUserHandler
	return u
}

// GetByIdRouter 更新用户的router
func (u *UserRouter) GetByIdRouter() Router {
	u.path = "/user/getById"
	u.handler = getByIdHandler
	return u
}

// UpdateUserRouter 更新用户的router
func (u *UserRouter) DeleteByIdRouter() Router {
	u.path = "/user/deleteById"
	u.handler = deleteByIdHandler
	return u
}

func deleteByIdHandler(resp http.ResponseWriter, req *http.Request) {
	// // 这里可以做一些数据合法性校验及安全校验，如os注入攻击
	userId, err := getIdFromReqHeader(req)
	if err != nil {
		// 应记录日志
		setParamErr(resp)
		return
	}
	err = controller.NewUserController(nil).DeleteUserById(userId)
	if err != nil {
		// 应记录日志
		setErrResp(resp, err.Error())
		return
	}
	setSuccessResp(resp, "succeed")
}

func getByIdHandler(resp http.ResponseWriter, req *http.Request) {
	// // 这里可以做一些数据合法性校验及安全校验，如os注入攻击
	userId, err := getIdFromReqHeader(req)
	if err != nil {
		// 应记录日志
		setParamErr(resp)
		return
	}
	userJsonStr, err := controller.NewUserController(nil).GetUserById(userId)
	if err != nil {
		// 应记录日志
		setErrResp(resp, err.Error())
		return
	}
	setSuccessResp(resp, userJsonStr)
}

// createUserHandler 创建用户的请求处理器，方法POST
func createUserHandler(resp http.ResponseWriter, req *http.Request) {
	// // 这里可以做一些数据合法性校验及安全校验，如os注入攻击
	user, err := getUserModelFromReqBody(req)
	if err != nil {
		// 应记录日志
		setParamErr(resp)
		return
	}
	err = controller.NewUserController(nil).CreateUser(user)
	if err != nil {
		// 应记录日志
		setErrResp(resp, err.Error())
		return
	}
	setSuccessResp(resp, "succeed")
}

// updateUserHandler 更新用户的请求处理器，方法POST
func updateUserHandler(resp http.ResponseWriter, req *http.Request) {
	// // 这里可以做一些数据合法性校验及安全校验，如os注入攻击
	user, err := getUserModelFromReqBody(req)
	if err != nil {
		// 应记录日志
		setParamErr(resp)
		return
	}
	err = controller.NewUserController(nil).UpdateUser(user)
	if err != nil {
		// 应记录日志
		setErrResp(resp, err.Error())
		return
	}
	setSuccessResp(resp, "succeed")
}

// Router 路由抽象类，解耦路由
type Router interface {
	GetPath() string
	GetHandler() http.HandlerFunc
}
