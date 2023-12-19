package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello" + request
	return nil
}

func main() {
	// 1.实例化一个服务
	listen, _ := net.Listen("tcp", ":3000")
	// 2.注册
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 3.启动服务
	for {
		conn, _ := listen.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
