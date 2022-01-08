package main

import (
	"net"
	"bufio"
	"fmt"
	"log"
	"io"
	"strings"
	"strconv"
	"errors"
)

func main() {
	Set("name", "matrix")
	Get("name")
}

type Client struct {
	net.Conn
	r *bufio.Reader
}

func Get(key string) {
	c, err := newClient()
	if err != nil {
		fmt.Println("client connect error: ", err)
		return
	}
	defer c.Close()

	klen := len(key)
	c.Write([]byte(fmt.Sprintf("G%d %s", klen, key)))
	resp, err := c.recvResponse()
	if err != nil {
		log.Println("recv reponse error: ", err)
		return
	}
	fmt.Println("resp: ", resp)
	c.Close()
}

func Set(key, val string) {
	c, err := newClient()
	if err != nil {
		fmt.Println("client connect error: ", err)
		return
	}
	defer c.Close()

	klen := len(key)
	vlen := len(val)
	fmt.Println("will send ", fmt.Sprintf("S%d %d %s%s", klen, vlen, key, val))
	c.Write([]byte(fmt.Sprintf("S%d %d %s%s", klen, vlen, key, val)))
	resp, err := c.recvResponse()
	if err != nil {
		log.Println("recv reponse error: ", err)
		return
	}
	fmt.Println("resp: ", resp)
}

func newClient() (*Client, error) {
	conn, err := net.Dial("tcp", "127.0.0.1:7641")
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	client := &Client{Conn: conn, r: bufio.NewReader(conn)}
	return client, nil
}

func (c *Client) recvResponse() (string, error) {
	vlen := readLen(c.r)
	if vlen == 0 {
		return "", nil
	}
	if vlen < 0 {
		err := make([]byte, -vlen)
		_, e := io.ReadFull(c.r, err)
		if e != nil {
			return "", e
		}
		return "", errors.New(string(err))
	}
	value := make([]byte, vlen)
	_, e := io.ReadFull(c.r, value)
	if e != nil {
		return "", e
	}
	return string(value), nil
}

func readLen(r *bufio.Reader) int {
	tmp, e := r.ReadString(' ')
	if e != nil {
		log.Println(e)
		return 0
	}
	l, e := strconv.Atoi(strings.TrimSpace(tmp))
	if e != nil {
		log.Println(tmp, e)
		return 0
	}
	return l
}
