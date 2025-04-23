package handler

import (
	"io"
	"net/http"

	"github.com/JoTaeYang/Admin/admin-back/service"
	"github.com/JoTaeYang/Admin/gpkg/pt"
	"github.com/gin-gonic/gin"
)

type ManagerHandler struct {
	service service.ManagerService
}

func NewManagerHandler(service service.ManagerService) *ManagerHandler {
	return &ManagerHandler{service: service}
}

func (h *ManagerHandler) GetManagerList(c *gin.Context) {
	users, err := h.service.Get(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h *ManagerHandler) PutManager(c *gin.Context) {
	errResponse := gin.H{"err": "invalid request"}
	req := &pt.ManagerCreateRequest{}
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	err = req.Unmarshal(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	users, err := h.service.Put(req.Id, req.Grade, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	_ = users

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
