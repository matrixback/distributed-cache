package tcp

import (
	"net"

	"github.com/matrixback/distributed-cache/cache"
)

type Server struct {
	cache.Cache
}

func (s *Server) Serve() {
	listen, err := net.Listen("tcp", ":7641")
	if err != nil {
		panic(err)
	}

	for {
		conn, e := listen.Accept()
		if e != nil {
			panic(e)
		}

		go s.handle(conn)
	}
}

func NewServer(c cache.Cache) *Server {
	return &Server{c}
}
