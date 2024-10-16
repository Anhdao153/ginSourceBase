package user

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")
	{
		user.GET("/:id/detail", UserDetail)
		user.POST("/", CreateUser)
		user.DELETE("/:id", UserDelete)
		user.PUT("/:id", UserUpdate)
	}
}
