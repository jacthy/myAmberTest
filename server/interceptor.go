package server

import "net/http"

// Handler 定义处理函数，同http处理函数handler
// 该函数用于真正处理grpc信号
type Handler http.HandlerFunc

// Interceptor 拦截器，该方法的设计特点：入参要有handler的入参以及handler，返回值要与handler相同
// info这种附带的值可以根据业务需要而定
type Interceptor func(resp http.ResponseWriter, req *http.Request, handler Handler)

// ChainInterceptor 拦截器调用链，生成最终链式拦截器
func ChainInterceptor(interceptorArr ...Interceptor) Interceptor {
	n := len(interceptorArr)

	return func(resp http.ResponseWriter, req *http.Request, handler Handler) {
		chainerFunc := func(interceptor Interceptor, handler Handler) Handler {
			return func(resp http.ResponseWriter, req *http.Request) {
				// 拦截器实际上在这里被装载进去，并被链式调用
				interceptor(resp, req, handler)
			}
		}
		chanHandler := handler
		// 倒序是为了顺序执行拦截器
		for i := n - 1; i >= 0; i-- {
			// 实际上的handler函数在这里被层层传递进去拦截器链中
			chanHandler = chainerFunc(interceptorArr[i], chanHandler)
		}
		chanHandler(resp, req)
	}
}
