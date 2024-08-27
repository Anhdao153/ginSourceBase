package prisma

import (
	"github/web-foreman/prisma/db"
)

var Prisma *db.PrismaClient

func run() (*db.PrismaClient, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	return client, nil
}

func PrismaInit() *db.PrismaClient {
	Client, err := run()
	if err != nil {
		panic(err)
	}
	Prisma = Client
	return Client
}
