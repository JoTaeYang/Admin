package handler

import (
	"io"
	"net/http"

	"github.com/JoTaeYang/Admin/admin-back/service"
	"github.com/JoTaeYang/Admin/gpkg/pt"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

type LoginHandler struct {
	service service.LoginService
}

func NewLoginHandler(service service.LoginService) *LoginHandler {
	return &LoginHandler{service: service}
}

func (h *LoginHandler) Login(c *gin.Context) {
	req := &pt.LoginRequest{}
	data, err := io.ReadAll(c.Request.Body)

	errResponse := gin.H{"err": "invalid request"}

	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	err = protojson.Unmarshal(data, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	token, err := h.service.Login(req.Id, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	res := &pt.LoginResponse{Token: token}
	c.JSON(http.StatusOK, res)
}
