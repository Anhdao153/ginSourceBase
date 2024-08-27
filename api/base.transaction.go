package api

import (
	"context"
	"github/web-foreman/prisma"
	"github/web-foreman/prisma/db"
)

func MultipleTransaction(commands []db.PrismaTransaction) error {
	ctx := context.Background()
	if err := prisma.Prisma.Prisma.Transaction(commands...).Exec(ctx); err != nil {
		return err
	}
	return nil
}
