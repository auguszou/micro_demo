package main

import (
	"fmt"
	"./hello_world"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

func main() {
	service := micro.NewService(
		micro.Name("hello_world"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init()

	greeter := hello_world.NewHelloWorldClient("hello_world", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &hello_world.HelloWorldRequest{Name: "Alice"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
