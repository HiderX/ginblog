package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins:  []string{"http://localhost:3000"},
			AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
			AllowHeaders:  []string{"Origin"},
			ExposeHeaders: []string{"Content-Length", "Authorization"},
			MaxAge:        12 * time.Hour,
		})
	}
}
