package helloworld

import (
	"context"
	"log"
)

type HelloWorldService struct {
	UnimplementedHelloWorldServer
}

func (hw *HelloWorldService) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &HelloReply{
		Message: "",
	}, nil
}
