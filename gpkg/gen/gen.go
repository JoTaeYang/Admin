package gen

import (
	"time"

	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/bwmarrin/snowflake"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	node *snowflake.Node
)

func UUID() string {
	return uuid.NewString()
}

func SnowFlake() int64 {
	if node == nil {
		node, _ = snowflake.NewNode(1)
	}

	return node.Generate().Int64()
}

func GenerateJWT(userID, secretKey string) (string, error) {
	// 클레임 설정
	claims := jwt.MapClaims{
		"sub": userID,                               // subject
		"exp": time.Now().Add(time.Hour * 1).Unix(), // 만료 시간
		"iat": time.Now().Unix(),                    // 발행 시간
		"iss": "my-app",                             // 발행자
	}

	// 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 서명된 토큰 문자열 반환
	tokenString, err := token.SignedString(converter.ZeroCopyStringToBytes(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
