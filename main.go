package main

import (
	env "github/web-foreman/env"
	"github/web-foreman/middleware"
	"github/web-foreman/prisma"
	"github/web-foreman/routes"
	"github/web-foreman/swagger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var Server *gin.Engine = gin.Default()

func main() {
	env.Use()
	gin.ForceConsoleColor()
	prisma.PrismaInit()
	routes.GinRoutes(Server)
	middleware.Middlewares(Server)
	Server.Static("/docs", "./docs")
	swagger.Use(Server)
	// listen and serve on 0.0.0.0:8080
	Server.Run(":" + viper.GetString("PORT"))

	// start grpc
	s := grpc.NewServer()
	routes.GrpcRoutes(s)
}
