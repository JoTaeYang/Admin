package service

import (
	"errors"
	"time"

	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/repo"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	Login(id string, pw string) (string, error)
}

type loginService struct {
	config *config.Configs
	loader *model.Loader
}

func NewLoginService(loader *model.Loader, config *config.Configs) LoginService {
	return &loginService{
		loader: loader,
		config: config,
	}
}

func (s *loginService) Login(id string, pw string) (string, error) {
	selector := model.NewSelector(id)

	selector.AddSingle(model.EManager, &repo.ManagerRepository{})

	db := bsql.RDB.GetAdminDB()

	results, err := s.loader.LoadTx(db, selector)
	if err != nil {
		return "", err
	}

	ctx := model.NewDataContext(results)

	user, ok := model.GetFromContext[*model.Manager](ctx, model.EManager)
	if !ok {
		return "", errors.New("not found datas")
	}

	pwBytes := converter.ZeroCopyStringToBytes(pw)
	dbPWBytes := converter.ZeroCopyStringToBytes(user.Password)

	err = bcrypt.CompareHashAndPassword(dbPWBytes, pwBytes)
	if err != nil {
		return "", err
	}

	return generateJWT(id, s.config.GetSecretKey())
}

func generateJWT(userID, secretKey string) (string, error) {
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
