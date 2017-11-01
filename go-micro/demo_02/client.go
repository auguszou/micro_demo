package main

import (
	sum "./sum"
	"fmt"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// micro query go.micro.fnc.sum Sum.CalSum '{"x": 1, "y": 10}'
func main() {
	f := micro.NewFunction(micro.Name("Sum"))

	c := sum.NewSumClient("go.micro.fnc.sum", f.Client())

	resp, err := c.CalSum(context.TODO(), &sum.SumRequest{X: 1, Y: 10})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Sum)
}
