package handler

import (
	"io"
	"net/http"
	"strings"

	"github.com/JoTaeYang/Admin/auth-back/service"
	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/JoTaeYang/Admin/gpkg/pt"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/jsonpb"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	errResponse := gin.H{"err": "invalid request"}
	req := &pt.AuthSignRequest{}
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, errResponse)
		return
	}

	err = jsonpb.Unmarshal(strings.NewReader(converter.ZeroCopyByteToString(data)), req)
	if err != nil {
		c.JSON(http.StatusOK, errResponse)
		return
	}

	err = h.service.SignUp(req.Uid)
	if err != nil {
		c.JSON(http.StatusOK, errResponse)
		return
	}

	//res := &pt.LoginResponse{Token: token}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
