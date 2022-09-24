package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err.Error())
	}
	var response string
	err = client.Call("HelloService.Hello", "lds", &response)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(response)
}
