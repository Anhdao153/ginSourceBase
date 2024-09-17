package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
}

func (v *Validator) nonNull() bool {
	return true
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		c.Next()
		// after request
		if c.Writer.Status() == 400 {
			if len(c.Errors) <= 0 {
				return
			}
			var errorMsgs []ErrorMsg
			for _, e := range c.Errors {
				validationErr, ok := e.Err.(validator.ValidationErrors)
				if !ok {
					return
				}
				for _, fe := range validationErr {
					errorMsg := ErrorMsg{
						Field:   fe.Field(),
						Message: getErrorMsg(fe),
					}
					errorMsgs = append(errorMsgs, errorMsg)
				}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errorMsgs})
		}
	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "email":
		return "Must be a valid email" + fe.Param()
	}
	return "Unknown error"
}
