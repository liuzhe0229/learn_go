package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1.建立连接
	client, err := rpc.Dial("tcp", ":3000")
	if err != nil {
		panic("Failed to connect！")
	}
	var reply *string = new(string)
	err = client.Call("HelloService.Hello", "zhe.liu", reply)
	if err != nil {
		panic("Failed to call the HelloService.Hello")
	}
	fmt.Println(*reply)
}
