package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 1.建立连接
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		panic("Failed to connect！")
	}
	var reply *string = new(string)
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "zhe.liu", reply)
	if err != nil {
		panic("Failed to call the HelloService.Hello")
	}
	fmt.Println(*reply)
}
