package middleware

import "github.com/gin-gonic/gin"

func Middlewares(gin *gin.Engine) {
	gin.Use(LoggerMiddleware())
	// validate
	gin.Use(ValidatorMiddleware())
}
