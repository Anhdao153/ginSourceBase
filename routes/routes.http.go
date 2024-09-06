package routes

import (
	"github/web-foreman/api/health"
	"github/web-foreman/api/user"

	"github.com/gin-gonic/gin"
)

func GinRoutes(gin *gin.Engine) {
	v1 := gin.Group("/api/v1")
	user.UserRoutes(v1)
	health.HealthRoutes(v1)
}
