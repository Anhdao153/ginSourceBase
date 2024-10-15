package main

import (
	env "github/web-foreman/env"
	"github/web-foreman/middleware"
	"github/web-foreman/prisma"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var Server *gin.Engine = gin.Default()

func main() {
	env.Use()
	gin.ForceConsoleColor()
	prisma.PrismaInit()
	defer func() {
		if err := prisma.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	grpc := grpc.NewServer()

	serverInit := ServerInit{
		grpc:    grpc,
		restApi: Server,
	}
	middleware.Middlewares(serverInit.restApi)
	serverInit.restApi.Static("/docs", "./docs")
	serverInit.RunGrpcAndRestApiOnPort("8080")
}
