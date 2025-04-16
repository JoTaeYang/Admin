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

	url := "/big/admin"

	LoginRouter := r.Group(url)
	{
		h := handler.NewLoginHandler(svc)
		LoginRouter.POST("/login", h.Login)
	}

	ManagementRouter := r.Group(url + "/management")
	{
		ManageManagerRouter := ManagementRouter.Group("/account")
		{
			svc := service.NewManagerService(loader, &cfg)
			h := handler.NewManagerHandler(svc)
			ManageManagerRouter.GET("/list", h.GetManagerList)
			ManageManagerRouter.POST("/create", h.PutManager)
		}

	}

	CheckRouter := r.Group(url)
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
