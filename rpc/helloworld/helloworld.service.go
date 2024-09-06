package helloworld

import "context"

type HelloWorldService struct {
	UnimplementedHelloWorldServer
}

func (hw *HelloWorldService) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return &HelloReply{
		Message: "",
	}, nil
}
