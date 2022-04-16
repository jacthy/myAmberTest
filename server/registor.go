package server

import (
	"fmt"
	"github.com/liaojuntao/server/router"
	"net/http"
)

const (
	routerNum      = 4 // 路由器初始化数量
	interceptorNum = 2 // 拦截器初始化数量
)

// loadRouter 加载路由
func loadRouter() []router.Router {
	r := make([]router.Router, 0, routerNum)
	return append(r,
		router.GetUserRouter().CreateUserRouter(),
	)
}

// serverInterceptors 设置拦截器,利用链式调用实现解耦，插件式添加拦截器
func serverInterceptors() []Interceptor {
	i := make([]Interceptor, 0, interceptorNum)
	return append(i,
		safeValidInterceptor,
	)
}

// defaultInterceptor
func defaultInterceptor(resp http.ResponseWriter, req *http.Request, handler Handler) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("oh no, fatal err: %v \n", err) // 这里可以做一个告警的入口，及时报警处理
		}
	}()
	handler(resp, req)
}

// safeValidInterceptor 参数安全校验器
func safeValidInterceptor(resp http.ResponseWriter, req *http.Request, handler Handler) {
	println("safeValidInterceptor")
	fmt.Printf("req header: %v\n", req)
	handler(resp, req)
}
