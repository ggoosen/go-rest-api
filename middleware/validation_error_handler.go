package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidationErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any binding errors
		bindingErrors := c.Errors.ByType(gin.ErrorTypeBind)
		if len(bindingErrors) > 0 {
			var validationErrors validator.ValidationErrors
			if err := bindingErrors[0].Err; err != nil {
				if errors.As(err, &validationErrors) {
					errorMessages := make(map[string]string)
					for _, fieldErr := range validationErrors {
						fieldName := fieldErr.Field()
						switch fieldErr.Tag() {
						case "required":
							errorMessages[fieldName] = fieldName + " is required"
						case "email":
							errorMessages[fieldName] = fieldName + " must be a valid email address"
						default:
							errorMessages[fieldName] = fieldName + " is invalid"
						}
					}
					c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
				} else {
					// If the error is not a validation error, return it as is
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				}
				return
			}
		}
	}
}
