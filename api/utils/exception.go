package utils

import "github.com/gin-gonic/gin"

func ThrowException(context *gin.Context, status int, err error) {
	context.Error(err)
	context.Status(status)
}
