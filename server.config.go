package main

import (
	"github/web-foreman/middleware"
	"github/web-foreman/routes"
	"github/web-foreman/swagger"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

type ServerInit struct {
	grpc    *grpc.Server
	restApi *gin.Engine
}

func (server *ServerInit) RunRestAPIOnly(port string) {
	routes.GinRoutes(server.restApi)
	middleware.Middlewares(server.restApi)
	server.restApi.Static("/docs", "./docs")
	server.restApi.Run(":" + port)
}

func (server *ServerInit) RunGrpcOnly(port string) {

	routes.GrpcRoutes(server.grpc)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server.grpc.Serve(lis)
}

func (server *ServerInit) RunGrpcAndRestApiOnPort(port string) {
	routes.GrpcRoutes(server.grpc)
	routes.GinRoutes(server.restApi)
	h2s := &http2.Server{}
	mixedHandler := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if it's a gRPC request, if so, serve using the gRPC server
		if r.ProtoMajor == 2 && r.Header.Get("Content-Type") == "application/grpc" {
			server.grpc.ServeHTTP(w, r)
		} else {
			// Otherwise, serve using the Gin HTTP server
			server.restApi.ServeHTTP(w, r)
		}
	}), h2s)
	serverPort := ":" + port

	if err := http.ListenAndServe(serverPort, mixedHandler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (server *ServerInit) RunGrpcAndRestApiDifferentPort(portRestAPI string, portGrpc string) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func() {
		defer waitGroup.Done()
		routes.GinRoutes(server.restApi)
		middleware.Middlewares(server.restApi)
		server.restApi.Static("/docs", "./docs")
		swagger.Use(server.restApi)
		server.restApi.Run(":" + portRestAPI)
	}()

	func() {
		defer waitGroup.Done()

		routes.GrpcRoutes(server.grpc)
		lis, err := net.Listen("tcp", ":"+portGrpc)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		server.grpc.Serve(lis)
	}()
	waitGroup.Wait()
}
