package handler

const HelloServiceName = "handler/HelloService"

type NewHelloService struct{}

func (s *NewHelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request + " !"
	return nil
}
