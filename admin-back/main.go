package main

import (
	"net/http"

	"github.com/JoTaeYang/Admin/admin-back/handler"
	"github.com/JoTaeYang/Admin/admin-back/service"
	mw "github.com/JoTaeYang/Admin/gpkg/middleware"
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
	r.Use(mw.AuthMiddleware(&cfg))

	loader := model.NewLoader()
	svc := service.NewLoginService(loader, &cfg)
	h := handler.NewLoginHandler(svc)

	LoginRouter := r.Group("/big/admin")
	{
		LoginRouter.POST("/login", h.Login)
	}

	CheckRouter := r.Group("/big/admin")
	{
		CheckRouter.GET("/me", func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})
	}

	return r
}

func main() {
	InitConfig()

	engine = InitRouter()

	engine.Run(cfg.Server.Port)
}
