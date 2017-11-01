package main

import (
	proto "./sum"
	// "fmt"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Sum struct {
}

func (s *Sum) CalSum(ctx context.Context, req *proto.SumRequest, resp *proto.SumResponse) error {
	resp.Sum = req.X + req.Y
	return nil
}

func main() {
	f := micro.NewFunction(micro.Name("go.micro.fnc.sum"))

	f.Init()

	f.Handle(new(Sum))

	f.Run()
}
