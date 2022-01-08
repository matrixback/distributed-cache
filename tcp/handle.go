package tcp

import (
	"bufio"
	"io"
	"log"
	"net"
)

func (s *Server) handle(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		op, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println("read conn error: ", err) 
			}
			return
		}

		if op == 'S' {
			err = s.set(conn, reader)
		} else if op == 'G' {
			err = s.get(conn, reader)
		} else if op == 'D' {
			err = s.del(conn, reader)
		} else {
			log.Println("invalid op: ", op)
		}

		if err != nil {
			log.Println("op error: ", err)
			return
		}
	}
}

func (s *Server) get(conn net.Conn, reader *bufio.Reader) error {
	key, err := s.readKey(reader)
	if err != nil {
		return err
	}
	val, err := s.Get(key)
	return sendResponse(val, err, conn)
}

func (s *Server) set(conn net.Conn, r *bufio.Reader) error {
	key, val, err := s.readKeyAndValue(r)
	if err != nil {
		return err
	}
	return sendResponse(nil, s.Set(key, val), conn)
}

func (s *Server) del(conn net.Conn, r *bufio.Reader) error {
	key, err := s.readKey(r)
	if err != nil {
		return err
	}
	return sendResponse(nil, s.Del(key), conn)
}