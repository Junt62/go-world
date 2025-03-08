package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

		c.Next()

		duration := time.Since(start)

		log.Printf("Response: %s %s in %v", c.Request.Method, c.Request.URL.Path, duration)
	}
}
