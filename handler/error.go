package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context, err error, statusCode ...int) {
	// Log the error (you can use a logging library here)
	fmt.Println("Error:", err)
	status := http.StatusInternalServerError
	if len(statusCode) > 0 {
		status = statusCode[0]
	}
	// Respond to the client with an error message
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email1":
		return "Should be an email"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}
