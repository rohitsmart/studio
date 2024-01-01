package util

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rohitsmart/studio/model"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func NewErrorResponse(c *gin.Context, errorDescription, userDescription, code string, statusCode int) {
	response := model.ErrorResponse{
		ErrorDescription: errorDescription,
		UserDescription:  userDescription,
		Code:             code,
	}
	c.JSON(statusCode, response)
}
