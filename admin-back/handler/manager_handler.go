package handler

import (
	"log"
	"net/http"

	"github.com/JoTaeYang/Admin/admin-back/service"
	"github.com/gin-gonic/gin"
)

type ManagerHandler struct {
	service service.LoginService
}

func NewManagerHandler(service service.LoginService) *ManagerHandler {
	return &ManagerHandler{service: service}
}

func (h *ManagerHandler) GetManagerList(c *gin.Context) {
	log.Println("Hello GetManager")
	//res := &pt.LoginResponse{Token: token}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
