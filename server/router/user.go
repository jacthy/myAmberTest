package router

import (
	"encoding/json"
	"github.com/liaojuntao/controller"
	"github.com/liaojuntao/infrastruct"
	"github.com/liaojuntao/infrastruct/repo"
	"io/ioutil"
	"net/http"
	"strconv"
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

// createUserHandler 创建用户的请求处理器，方法POST
func createUserHandler(resp http.ResponseWriter, req *http.Request) {
	// 这里可以做一些数据合法性校验，如os注入攻击
	user, err := getUserModelFromReqBody(req)
	if err != nil {
		setParamErr(resp)
		return
	}
	err = controller.NewUserController(repo.GetUserRepo()).CreateUser(user)
	if err != nil {
		setErrResp(resp, err.Error())
		return
	}
	setSuccessResp(resp, "succeed")
}

func getUserModelFromReqBody(req *http.Request) (*infrastruct.User, error) {
	userModel := new(infrastruct.User)
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(result, userModel); err != nil {
		return nil, err
	}
	return userModel, nil
}

func getIdFromReqHeader(req *http.Request) (int, error) {
	return strconv.Atoi(req.URL.Query().Get("userId"))
}

// Router 路由抽象类，解耦路由
type Router interface {
	GetPath() string
	GetHandler() http.HandlerFunc
}
