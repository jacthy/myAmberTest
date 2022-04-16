package server

import (
	"github.com/liaojuntao/server/router"
	"net/http"
)

type Server struct {
	addr string
	router []router.Router
	serverMux HttpServeMux
}

func NewServer(addr string) *Server {
	s := &Server{
		addr: addr,
	}
	return s
}

func (s *Server) init() {

}

type HttpServeMux struct {
	http.ServeMux
}

func (hsm *HttpServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hsm.ServeMux.ServeHTTP(w,r)
}



