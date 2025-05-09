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

type ShopHandler struct {
	service service.ShopService
}

func NewShopHandler(service service.ShopService) *ShopHandler {
	return &ShopHandler{service: service}
}

func (h *ShopHandler) Gacha(c *gin.Context) {
	errResponse := gin.H{"err": "invalid request"}
	req := &pt.ShopGachaRequest{}
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

	dataCtx, err := h.service.Gacha(c, req.Key)
	if err != nil {
		c.JSON(http.StatusOK, errResponse)
		return
	}

	res := &pt.ShopGachaResponse{Datas: api.MakeLoadResponse(dataCtx)}
	c.JSON(http.StatusOK, res)
}
