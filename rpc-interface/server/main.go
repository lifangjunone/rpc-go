package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"rpc-go/rpc-interface/service"
)

var _ service.HelloService = (*HelloService)(nil)

type HelloService struct{}

func (h *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello %s", request)
	return nil
}

func main() {
	// register helloService to rpc
	rpc.RegisterName(service.SERVICE_NAME, &HelloService{})
	lister, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err.Error())
	}
	for {
		conn, err := lister.Accept()
		if err != nil {
			log.Fatalf("accept connection is error %s", err.Error())
		}
		go rpc.ServeConn(conn)
	}
}
