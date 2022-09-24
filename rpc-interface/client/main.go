package main

import (
	"fmt"
	"net/rpc"
	"rpc-go/rpc-interface/service"
)

var _ service.HelloService = (*HelloServiceClient)(nil)

type HelloServiceClient struct {
	client *rpc.Client
}

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{
		client: client,
	}, nil
}

func (h *HelloServiceClient) Hello(request string, response *string) error {
	err := h.client.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), request, response)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	c, err := NewHelloServiceClient("tcp", ":1234")
	if err != nil {
		panic(err.Error())
	}
	var response string
	c.Hello("ldd", &response)
	fmt.Println(response)
}
