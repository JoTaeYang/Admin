package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const secretKey = "local_test_jwt"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err_code": "error",
			})
			c.Abort()
			return
		}

		prefix := "Bearer "
		token := authorization

		if strings.HasPrefix(authorization, prefix) {
			token = authorization[len(prefix):]
		}

		_ = token
	}
}
