package server

import (
	"github.com/liaojuntao/server/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Server 自己实现的一个web server
// 大的工程应用还是应该考虑用gin,beego等完善的框架提升开发效率
type Server struct {
	addr   string
	router []router.Router
	*http.ServeMux
	Interceptors []Interceptor
	log.Logger   // todo 日志模块，实现日志追踪，请求链追踪等
}

// NewServer 获取server实例
func NewServer(addr string) *Server {
	s := &Server{
		addr: addr,
	}
	return s
}

// Init server的初始化
func (s *Server) Init() {
	s.ServeMux = http.DefaultServeMux
	s.router = loadRouter()
	s.Interceptors = serverInterceptors()
}

// ServeHTTP 重写http.ServeMux的serveHttp方法，嵌入拦截器
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	finalInterceptor := ChainInterceptor(s.Interceptors...)
	finalInterceptor(w, r, s.ServeMux.ServeHTTP)
}

// Run 启动服务
func (s *Server) Run() error {
	errChan := make(chan error)
	stopChan := make(chan os.Signal) // 利用该stopChan可以做优雅退出
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	// 加载路由器
	for _, r := range s.router {
		println(r.GetPath())
		http.HandleFunc(r.GetPath(), r.GetHandler())
	}
	go func() {
		errChan <- http.ListenAndServe(s.addr, s)
	}()

	select {
	case err := <-errChan:
		return err
	case <-stopChan:
	}
	println("graceful stop：you had terminated this proc")
	return nil
}
