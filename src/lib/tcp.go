package lib


import (
	"errors"
	"fmt"
	"net"
	"time"
)


func TcpConnect(ip string, port uint16) (net.Conn, error) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	return net.Dial("tcp", addr)
}

func TcpAsyncConnect(ip string, port uint16, cb func(net.Conn, error)) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	go func() {
		conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
		cb(conn, err)
	}()
}

func Send(sock_h net.Conn, data []byte) error {
	if sock_h == nil {
		return errors.New("net handle nil")
	}

	buffSize := len(data)
	for {
		if buffSize <= 0 {
			break
		}

		n, err := sock_h.Write(data)
		if err != nil {
			return err
		}

		buffSize -= n
	}

	return nil
}

func Recv(sock_h net.Conn, data []byte) (int, error) {
	if sock_h == nil {
		return 0, errors.New("net handle nil")
	}
	return sock_h.Read(data)
}

func Close(sock_h net.Conn) error {
	if sock_h == nil {
		return nil
	}

	return sock_h.Close()
}
