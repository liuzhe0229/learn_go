package server_proxy

import (
	handler "learn_go/helloworld_rpc_new/handler"
	"net/rpc"
)

// 用自己都语音来理解就是这里只定义一个框架 也就是这里只表示 我需要有什么函数
// 至于函数都具体逻辑是什么 我不管，我只需要规定你传给RegisterHelloService都参数满足我定义的接口
// 这里的接口就是只需要你有一个Hello方法可以了
// 鸭子类型，你传进来的struct有Hello 就认为满足接口要求

type HelloServiceEr interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServiceEr) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
