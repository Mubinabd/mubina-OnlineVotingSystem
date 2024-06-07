package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	name := c.GetHeader("name")
	date := c.GetHeader("date")
	if name != "name" || date != "date" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	timestamp := time.Now()
	path := c.Request.URL.Path
	log.Printf("Authenticated request at %s for path %s\n", timestamp.Format(time.RFC3339), path)
	c.Next()
}
