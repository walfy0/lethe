package main

import (
	"github.com/lethe/config"
	"github.com/lethe/handler"
	"github.com/lethe/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thinkerou/favicon"
)

func init() {
	logrus.SetLevel(logrus.TraceLevel)
	config.InitMysql()
}

func main() {
	r := gin.New()
	r.Use(favicon.New("./favicon.ico"))
	r.Use(gin.Logger(),
		gin.Recovery(),
		middleware.LoginMiddleWare())
	v := r.Group("/lethe")
	{
		v.GET("ping", handler.Test)
		//auth
		v.POST("register", handler.Register)
		v.POST("send_mail", handler.SendEmail)
		v.POST("login", handler.Login)
		v.GET("logout", handler.Logout)
		//doc
		v.POST("doc/list", handler.DocList)
		v.POST("doc/update", handler.DocUpdate)
	}
	r.Run()
}
