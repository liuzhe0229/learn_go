package main

import (
	"fmt"
	"google.golang.org/grpc"
	"learn_go/stream_grpc_test/proto"
	"net"
	"sync"
	"time"
)

type Service struct {
}

func (s *Service) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("Send a message from server by single way%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		i++
		if i > 10 {
			break
		}
	}
	return nil
}
func (s *Service) PutStream(clientStream proto.Greeter_PutStreamServer) error {
	for {
		if a, err := clientStream.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}

	return nil
}
func (s *Service) AllStream(allStream proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		i := 1
		for {
			_ = allStream.Send(&proto.StreamResData{
				Data: fmt.Sprintf("Send a message for Server by two-way%v", time.Now().Unix()),
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
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		panic("Failed to listen the port 8000;" + err.Error())
	}

	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Service{})

	err = g.Serve(lis)
}
