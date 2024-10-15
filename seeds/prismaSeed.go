package main

import (
	"fmt"
	"github/web-foreman/api"
	"github/web-foreman/prisma"
	"github/web-foreman/prisma/db"

	"github.com/steebchen/prisma-client-go/runtime/transaction"
)

func main() {
	client := prisma.PrismaInit()
	transactionRoles := []transaction.Transaction{}
	// Create default role
	for _, roleName := range []db.RoleName{db.RoleNameAdmin, db.RoleNameManager, db.RoleNameUser} {
		role := client.Roles.CreateOne(
			db.Roles.RoleName.Set(roleName),
		).Tx()
		transactionRoles = append(transactionRoles, role)
	}
	api.MultipleTransaction(transactionRoles)
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(fmt.Errorf("could not disconnect: %w", err))
		}
	}()
}
