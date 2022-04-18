package server

import (
	"fmt"
	"github.com/liaojuntao/server/router"
	"net/http"
)

const (
	routerNum      = 4 // 路由器初始化数量
	interceptorNum = 4 // 拦截器初始化数量
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

// serverInterceptors 设置拦截器,利用链式调用实现解耦，插件式添加拦截器，如链路追踪器/权限控制器/安全校验器等
func serverInterceptors() []Interceptor {
	i := make([]Interceptor, 0, interceptorNum)
	return append(i,
		defaultInterceptor,
		safeValidInterceptor,
		tracerInterceptor,
		authInterceptor,
	)
}

func authInterceptor(resp http.ResponseWriter, req *http.Request, handler Handler) {
	// 这里可以进行鉴权/认证
	// 通过req的token，校验是否有权限，没有权限则要求重新登陆
	// 接口的权限控制可以通过权限码的方式，根据cookie取出用户权限码（前端通过开放权限码查询接口获取权限码然后传入）
	// 权限码对应的接口权限可以记录在权限表里，根据接口名和权限码作为联合索引，存在记录就表明有接口权限
	handler(resp,req)
}

func tracerInterceptor(resp http.ResponseWriter, req *http.Request, handler Handler) {
	// 这里可以实现链路追踪，记录req信息，为request打上一个id标签然后输入到日志文件中
	// 记录信息可以有：用户id，链路id（没有则创建，如uuid），时间，请求参数，响应参数，响应时间
	// 存储方案可以接通一些云存储，如阿里云，贾维斯
	// 也可以自己实现日志存储，但需要考虑空间，过期处理，检索，

	handler(resp,req)
}

// defaultInterceptor 默认拦截器
func defaultInterceptor(resp http.ResponseWriter, req *http.Request, handler Handler) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("oh no, fatal err: %v \n", err)
			_,writeErr:=resp.Write([]byte(systemErrStr))
			if writeErr != nil {
				// 记录日志
				println("写入响应失败:",writeErr.Error())
			}
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
