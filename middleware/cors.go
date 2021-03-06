package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cors.New(cors.Config{
			AllowAllOrigins: true,
			//AllowOrigins: []string{"*"},
			AllowMethods: []string{"*"},
			AllowHeaders: []string{"Origin"},
			ExposeHeaders: []string{"Content-Length", "Authorization"},
			AllowCredentials: false,
			MaxAge: 12 * time.Hour,
		})
	}
}
