package main

import (
	proto "./proto"
	"fmt"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctxt context.Context, req *proto.HelloRequest, resp *proto.HelloResponse) error {
	resp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
