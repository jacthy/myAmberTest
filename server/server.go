package server

import (
	"github.com/liaojuntao/server/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	addr   string
	router []router.Router
	*http.ServeMux
	Interceptors []Interceptor
	log.Logger // todo 日志模块，实现日志追踪，可以做请求链追踪等
}

func NewServer(addr string) *Server {
	s := &Server{
		addr: addr,
	}
	return s
}

func (s *Server) Init() {
	s.ServeMux = http.DefaultServeMux
	s.router = loadRouter()
	s.Interceptors = serverInterceptors()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	finalInterceptor := ChainInterceptor(s.Interceptors...)
	finalInterceptor(w, r, s.ServeMux.ServeHTTP)
}

func (s *Server) Run() error {
	errChan := make(chan error)
	stopChan := make(chan os.Signal) // 利用该stopChan可以做优雅退出
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	// 加载路由器
	for _, r := range s.router {
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
	println("graceful stop") // todo graceful stop
	return nil
}
