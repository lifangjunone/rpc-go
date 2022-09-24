package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello: %s", request)
	return nil
}

func main() {
	// 把RPC对外暴露的对象注册到RPC框架内部
	rpc.RegisterName("HelloService", &HelloService{})
	// Build tcp listen
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("listen is error %s", err.Error())
	}
	for {

		conn, err := listener.Accept()
		if err != nil {
			panic(err.Error())
		}
		go rpc.ServeConn(conn)
	}

}
