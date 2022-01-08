package tcp

import (
	"fmt"
	"net"
)

func sendResponse(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errString := err.Error()
		tmp := fmt.Sprintf("-%d ", len(errString)) + errString
		_, e := conn.Write([]byte(tmp))
		return e
	}
	vlen := fmt.Sprintf("%d ", len(value))
	fmt.Println("send:", append([]byte(vlen), value...),"end")
	_, e := conn.Write(append([]byte(vlen), value...))
	return e
}
