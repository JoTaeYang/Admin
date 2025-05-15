package handler

import (
	"io"
	"net/http"
	"strings"

	"github.com/JoTaeYang/Admin/game-back/service"
	"github.com/JoTaeYang/Admin/gpkg/api"
	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/JoTaeYang/Admin/gpkg/pt"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/jsonpb"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Load(c *gin.Context) {
	errResponse := gin.H{"err": "invalid request"}
	req := &pt.LoadRequest{}
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	err = jsonpb.Unmarshal(strings.NewReader(converter.ZeroCopyByteToString(data)), req)
	if err != nil {
		c.JSON(http.StatusOK, errResponse)
		return
	}

	dataCtx, err := h.service.Load(c, req.Uid)
	if err != nil {
		c.JSON(http.StatusOK, errResponse)
		return
	}

	res := &pt.LoadResponse{Datas: api.MakeLoadResponse(dataCtx)}
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) New(c *gin.Context) {
	errResponse := gin.H{"err": "invalid request"}
	req := &pt.NewUserRequest{}
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	err = jsonpb.Unmarshal(strings.NewReader(converter.ZeroCopyByteToString(data)), req)
	if err != nil {
		c.JSON(http.StatusOK, errResponse)
		return
	}

	if req.Id == "" || req.Name == "" {
		c.JSON(http.StatusOK, errResponse)
		return
	}

	err = h.service.New(c, req.Id, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	//res := &pt.LoginResponse{Token: token}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
