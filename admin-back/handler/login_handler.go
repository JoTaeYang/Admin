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
	errResponse := gin.H{"err": "invalid request"}
	req := &pt.LoginRequest{}
	data, err := io.ReadAll(c.Request.Body)
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

	c.SetCookie(
		"token",     // 쿠키 이름
		token,       // 쿠키 값 (JWT)
		3600,        // 유효 시간(초) → 1시간
		"/",         // 경로
		"localhost", // 도메인 (로컬이면 "localhost")
		false,       // Secure: HTTPS에서만 전송됨
		true,        // HttpOnly: JS에서 접근 불가
	)

	//res := &pt.LoginResponse{Token: token}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
