package handler

import (
	"log"
	"net/http"

	"github.com/JoTaeYang/Admin/admin-back/service"
	"github.com/gin-gonic/gin"
)

type ManagerHandler struct {
	service service.ManagerService
}

func NewManagerHandler(service service.ManagerService) *ManagerHandler {
	return &ManagerHandler{service: service}
}

func (h *ManagerHandler) GetManagerList(c *gin.Context) {
	log.Println("Hello GetManager")

	err := h.service.Get()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
