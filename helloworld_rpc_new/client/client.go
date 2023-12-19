package main

import (
	"fmt"
	"learn_go/helloworld_rpc_new/client_proxy"
)

func main() {
	// 1.建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:3000")
	var reply *string = new(string)
	// 2.远程调用
	err := client.Hello("zhe.liu", reply)
	if err != nil {
		panic("Failed to call the HelloService.Hello")
	}
	fmt.Println(*reply)
}
