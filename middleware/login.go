package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lethe/common"
)

var ignoreURL = []string{
	"/lethe/login",
	"/lethe/register",
	"/lethe/send_mail",
}

func LoginMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, url := range ignoreURL {
			if url == c.Request.URL.Path {
				c.Next()
				return
			}
		}
		_, err := c.Cookie(common.UserId)
		if err != nil {
			c.AbortWithStatus(400)
			return
		}
		c.Next()
	}
}
