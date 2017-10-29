package main

import (
	"fmt"
	"./hello_world"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type HelloWorld struct{}

func (g *HelloWorld) Hello(ctx context.Context, req *hello_world.HelloWorldRequest, rsp *hello_world.HelloWorldResponse) error {
	rsp.Greeting = "Hello World: " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("hello_world"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init()

	hello_world.RegisterHelloWorldHandler(service.Server(), new(HelloWorld))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
