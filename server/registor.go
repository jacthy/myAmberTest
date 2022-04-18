package server

import (
	"fmt"
	"github.com/liaojuntao/server/router"
	"net/http"
)

const (
	routerNum      = 4 // 路由器初始化数量
	interceptorNum = 2 // 拦截器初始化数量
	systemErrStr   = "{\"errCode\":5000,\"errMessage\":\"系统内部错误\"}"
)

// loadRouter 加载路由
func loadRouter() []router.Router {
	r := make([]router.Router, 0, routerNum)
	return append(r,
		router.CreateUserRouter(),
		router.UpdateUserRouter(),
		router.GetByIdRouter(),
		router.DeleteByIdRouter(),
	)
}

// serverInterceptors 设置拦截器,利用链式调用实现解耦，插件式添加拦截器
func serverInterceptors() []Interceptor {
	i := make([]Interceptor, 0, interceptorNum)
	return append(i,
		defaultInterceptor,
		safeValidInterceptor,
	)
}

// defaultInterceptor 默认拦截器
func defaultInterceptor(resp http.ResponseWriter, req *http.Request, handler Handler) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("oh no, fatal err: %v \n", err)
			resp.Write([]byte(systemErrStr))
			resp.WriteHeader(http.StatusInternalServerError)
		}
	}()
	handler(resp, req)
}

// safeValidInterceptor 参数安全校验器
func safeValidInterceptor(resp http.ResponseWriter, req *http.Request, handler Handler) {
	println("safeValidInterceptor")
	// 这里也可以做一些数据合法性校验，如所有入参的os注入攻击
	handler(resp, req)
}
