package tcp

/*
缓存的 tcp 服务实现，服务端与客户端采用 ABNF 协议通信

command = op key | key-value
op = 'S' | 'G' | 'D'
key = bytes-array
bytes-array = lenght SP content
length = 1 * DIGIT
content = *OCTET
key-value = length SP length SP content content
response = error | bytes-array
error = '-' bytes-array

DIGIT 取值 0~9
OCTET 取值 0x00~0xFF
*/

import (
	"net"
	"log"

	"github.com/matrixback/distributed-cache/cache"
)

type Server struct {
	cache.Cache
}

func (s *Server) Serve() {
	log.Println("cache server is starting, port: 7641")
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
