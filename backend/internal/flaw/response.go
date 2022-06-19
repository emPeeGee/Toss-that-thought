package flaw

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO: For future https://stackoverflow.com/questions/69948784/how-to-handle-errors-in-gin-middleware

func Error(c *gin.Context, code int, message, details string) {
	c.AbortWithStatusJSON(code, map[string]interface{}{
		"message": message,
		"details": details,
	})
}

func InternalServer(c *gin.Context, message, details string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
		"message": message,
		"details": details,
	})
}

func BadRequest(c *gin.Context, message, details string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
		"message": message,
		"details": details,
	})
}

func NotFound(c *gin.Context, message, details string) {
	c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
		"message": message,
		"details": details,
	})
}

// TODO: For future
func Unauthorized(c *gin.Context, message, details string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
		"message": message,
		"details": details,
	})
}
