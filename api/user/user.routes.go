package user

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")
	{
		user.GET("/detail", UserDetail)
		user.POST("/", CreateUser)
	}
}
