package server

import (
	"github.com/liaojuntao/server/router"
)

const (
	routerNum = 4
)

func (s *Server) loadRouter() {
	r := make([]router.Router,0,routerNum)
	s.router = append(r,
		router.GetUserRouter().CreateUserRout(),
		)
}
