package routes

import (
	"context"
	"github/web-foreman/rpc/helloworld"
	"log"

	"google.golang.org/grpc"
)

type server struct {
	helloworld.UnimplementedHelloWorldServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
func GrpcRoutes(grpc *grpc.Server) {
	helloworld.RegisterHelloWorldServer(grpc, &server{})
}
