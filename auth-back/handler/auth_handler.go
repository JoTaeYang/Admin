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

func (h *AuthHandler) Login(c *gin.Context) {
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

	token, err := h.service.Login(req.Uid)
	if err != nil {
		c.JSON(http.StatusOK, errResponse)
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

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
