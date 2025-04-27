package service

import (
	"errors"

	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/JoTaeYang/Admin/gpkg/gen"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/repo"
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

	return gen.GenerateJWT(id, s.config.GetSecretKey())
}
