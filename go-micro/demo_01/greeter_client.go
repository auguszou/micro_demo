package main

import (
	proto "./proto"
	"fmt"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
	)

	greeter := proto.NewGreeterClient("greeter", service.Client())

	resp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Alice"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Greeting)
}
