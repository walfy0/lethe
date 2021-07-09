package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var ignoreURL = []string{
	"/lethe/login",
}

func LoginMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, url := range ignoreURL {
			if url == c.Request.URL.Path {
				c.Next()
				return
			}
		}
		cookie, err := c.Cookie("UserInfo")
		if err != nil {
			c.AbortWithStatus(400)
			return
		}
		logrus.Info(cookie)
		c.Next()
	}
}
