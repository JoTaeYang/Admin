package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "local_test_jwt"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		skipPaths := map[string]bool{
			"/login": true,
		}

		if skipPaths[c.FullPath()] {
			c.Next()
			return
		}

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
		validToken, err := validateJWT(token)
		if err != nil || !validToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := validToken.Claims.(jwt.MapClaims); ok {
			c.Set("userID", claims["sub"])
		}

		c.Next()
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 서명 알고리즘 체크 (선택사항이지만 보안상 추천)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
}
