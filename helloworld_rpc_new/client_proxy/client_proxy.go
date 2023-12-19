package client_proxy

import (
	"learn_go/helloworld_rpc_new/handler"
	"net/rpc"
)

// 相当于在这个里面把所有的逻辑都聚合在 HellServiceStub这个struct 上面
// client 要做都也就是 New一个（Go都结构体是没有实例化都）
// new 一个实例后 后续都逻辑都挂在这个实例上

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protocol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protocol, address)
	if err != nil {
		panic("Failed to connected Service!")
	}
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
