package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(cfg *config.Configs) gin.HandlerFunc {
	return func(c *gin.Context) {
		skipPaths := map[string]bool{
			"login": true,
		}

		path := c.FullPath()
		for _, v := range strings.Split(path, "/") {
			if skipPaths[v] {
				c.Next()
				return
			}
		}

		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}
		secretKey := converter.ZeroCopyStringToBytes(cfg.GetSecretKey())
		validToken, err := validateJWT(tokenString, secretKey)
		if err != nil || !validToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := validToken.Claims.(jwt.MapClaims); ok {
			c.Set("userID", claims["sub"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func validateJWT(tokenString string, secretKey []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 서명 알고리즘 체크 (선택사항이지만 보안상 추천)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
}
