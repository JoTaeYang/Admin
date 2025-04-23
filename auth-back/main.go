package main

import (
	"github.com/JoTaeYang/Admin/auth-back/handler"
	"github.com/JoTaeYang/Admin/auth-back/service"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func CORSM() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
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

	loader := model.NewLoader()
	svc := service.NewAuthService(loader, &cfg)
	url := "/big/auth"

	AuthRouter := r.Group(url)
	{
		h := handler.NewAuthHandler(svc)
		AuthRouter.POST("/sign", h.SignUp)
	}

	return r
}

func main() {
	InitConfig()

	engine = InitRouter()

	engine.Run(cfg.Server.Port)
}
