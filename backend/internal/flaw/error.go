package flaw

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func Error(c *gin.Context, code int, message, details string) {
	c.AbortWithError(code, ErrorResponse{Status: code, Message: message, Details: details})
}

func InternalServer(c *gin.Context, message, details string) {
	c.AbortWithError(http.StatusInternalServerError, ErrorResponse{Status: http.StatusInternalServerError, Message: message, Details: details})
}

func BadRequest(c *gin.Context, message, details string) {
	c.AbortWithError(http.StatusBadRequest, ErrorResponse{Status: http.StatusBadRequest, Message: message, Details: details})
}

func NotFound(c *gin.Context, message, details string) {
	c.AbortWithError(http.StatusNotFound, ErrorResponse{Status: http.StatusNotFound, Message: message, Details: details})
}

// TODO: For future
func Unauthorized(c *gin.Context, message, details string) {
	c.AbortWithError(http.StatusUnauthorized, ErrorResponse{Status: http.StatusUnauthorized, Message: message, Details: details})
}
