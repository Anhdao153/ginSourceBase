package user

import (
	"github/web-foreman/prisma"
	"github/web-foreman/prisma/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserDetail(c *gin.Context) {
	res := UserResponse{
		Id:   1234,
		Name: "1234",
	}
	c.JSON(http.StatusOK, &res)

}

func CreateUser(c *gin.Context) {
	// ctx := context.Background()
	user := prisma.Prisma.Users.CreateOne(
		db.Users.Email.Set("example.com"),
		db.Users.Password.Set(true),
		db.Users.Username.Set("Hell"),
	)
	// user.Exec(ctx)
	user.Tx()
	c.JSON(http.StatusCreated, &user)
}
