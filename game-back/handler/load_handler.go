package handler

import (
	"io"
	"net/http"

	"github.com/JoTaeYang/Admin/game-back/service"
	"github.com/JoTaeYang/Admin/gpkg/pt"
	"github.com/gin-gonic/gin"
)

type LoadHandler struct {
	service service.LoadService
}

func NewLoadHandler(service service.LoadService) *LoadHandler {
	return &LoadHandler{service: service}
}

func (h *LoadHandler) Load(c *gin.Context) {
	errResponse := gin.H{"err": "invalid request"}
	req := &pt.LoadRequest{}
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

	//res := &pt.LoginResponse{Token: token}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
