package main

import (
	"github.com/JoTaeYang/Admin/admin-back/handler"
	"github.com/JoTaeYang/Admin/admin-back/service"
	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func CORSM() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Method", "GET, DELETE, POST")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	}
}

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(CORSM())

	svc := service.NewLoginService()
	h := handler.NewLoginHandler(svc)

	LoginRouter := r.Group("/big/admin")
	{
		LoginRouter.POST("/login", h.Login)
	}

	return r
}

func main() {
	InitConfig()

	engine = InitRouter()

	engine.Run(cfg.Server.Port)
}
