package main

import (
	"learn_go/helloworld_rpc_new/handler"
	"learn_go/helloworld_rpc_new/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	// 1.实例化一个服务
	listen, _ := net.Listen("tcp", ":3000")
	// 2.注册
	_ = server_proxy.RegisterHelloService(&handler.NewHelloService{})
	// 3.启动服务
	for {
		conn, _ := listen.Accept()
		go rpc.ServeConn(conn)
	}

}
