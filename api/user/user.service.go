package user

import (
	"context"
	"github/web-foreman/api/utils"
	"github/web-foreman/prisma"
	"github/web-foreman/prisma/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserDetail(c *gin.Context) {
	ctx := context.Background()
	userReqDetail := UserDetailRequest{}
	if err := c.BindUri(&userReqDetail); err != nil {
		utils.ThrowException(c, http.StatusBadRequest, err)
		return
	}

	detail, err := prisma.Prisma.Users.FindFirst(db.Users.ID.Equals(userReqDetail.Id)).With(db.Users.Role.Fetch()).Exec(ctx)
	if err != nil {
		utils.ThrowException(c, http.StatusNotFound, err)
		return
	}

	res := UserResponse{
		Username: detail.Username,
		Id:       detail.ID,
		Email:    detail.Email,
		RoleName: string(detail.Role().RoleName),
	}

	c.JSON(http.StatusOK, &res)
}

func UserDelete(c *gin.Context) {
	ctx := context.Background()
	userReqDetail := UserDetailRequest{}
	if err := c.BindUri(&userReqDetail); err != nil {
		utils.ThrowException(c, http.StatusBadRequest, err)
		return
	}

	res, err := prisma.Prisma.Users.FindUnique(db.Users.ID.Equals(userReqDetail.Id)).Delete().Exec(ctx)
	if err != nil {
		utils.ThrowException(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, &res)
}

func CreateUser(c *gin.Context) {
	ctx := context.Background()
	userReq := UserCreateRequest{}
	if err := c.ShouldBindJSON(&userReq); err != nil {
		utils.ThrowException(c, http.StatusBadRequest, err)
		return
	}

	user := prisma.Prisma.Users.CreateOne(
		db.Users.Email.Set(userReq.Email),
		db.Users.Password.Set(userReq.Password),
		db.Users.Username.Set(userReq.Username),
		db.Users.Role.Link(db.Roles.RoleName.Equals(db.RoleNameUser)),
	)
	user.Exec(ctx)
	c.JSON(http.StatusCreated, &user)
}

func UserUpdate(c *gin.Context) {
	ctx := context.Background()
	userReq := UserUpdateRequest{}
	userId := UserDetailRequest{}
	err := c.ShouldBindJSON(&userReq)
	err2 := c.BindUri(&userId)
	if err != nil || err2 != nil {
		utils.ThrowException(c, http.StatusBadRequest, err)
		return
	}

	user := prisma.Prisma.Users.FindUnique(db.Users.ID.Equals(userId.Id)).Update(
		db.Users.Email.Set(userReq.Email),
		db.Users.Password.Set(userReq.Password),
		db.Users.Username.Set(userReq.Username),
	)
	user.Exec(ctx)
	c.JSON(http.StatusCreated, &user)
}
