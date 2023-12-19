package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learn_go/stream_grpc_test/proto"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		panic("Failed to connect the server")
	}

	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	// 1.服务端流模式
	go func() {
		defer wg.Done()
		res, err := c.GetStream(context.Background(), &proto.StreamReqData{Data: "Send a request from Zhe.liu"})
		for {
			a, err := res.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(a)
		}

		if err != nil {
			panic("Send message Failed!")
		}
	}()

	// 2.客户端流模式
	go func() {
		defer wg.Done()
		putSteam, _ := c.PutStream(context.Background())
		i := 0
		for {
			_ = putSteam.Send(&proto.StreamReqData{
				Data: fmt.Sprintf("Send a message from Client by Signle way %v", time.Now().Unix()),
			})
			time.Sleep(time.Second)
			i++
			if i > 10 {
				break
			}
		}
	}()

	// 3双向流模式
	allStream, _ := c.AllStream(context.Background())

	go func() {
		defer wg.Done()
		i := 1
		for {
			_ = allStream.Send(&proto.StreamReqData{
				Data: fmt.Sprintf("Send a message for Client by tow-way%v", time.Now().Unix()),
			})
			time.Sleep(time.Second)
			i++
			if i > 10 {
				break
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if a, err := allStream.Recv(); err != nil {
				fmt.Println(err)
				break
			} else {
				fmt.Println(a.Data)
			}
		}
	}()

	wg.Wait()
}
