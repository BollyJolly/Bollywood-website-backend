package middleware

import (
	"github.com/gin-gonic/gin"
)

// Auth is a placeholder for future JWT authentication.
// For now it is a no-op so routes can be registered against it later.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: validate JWT / session and set userID on the context
		c.Next()
	}
}
