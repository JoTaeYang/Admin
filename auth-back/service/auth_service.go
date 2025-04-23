package service

import (
	"errors"

	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/repo"
)

type AuthService interface {
	SignUp(id string) error
}

type authService struct {
	config *config.Configs
	loader *model.Loader
}

func NewAuthService(loader *model.Loader, config *config.Configs) AuthService {
	return &authService{
		loader: loader,
		config: config,
	}
}

func (s *authService) SignUp(id string) error {
	selector := model.NewSelector(id)

	selector.AddSingle(model.EIdentity, &repo.IdentityRepository{})

	db := bsql.RDB.GetIdentityDB()
	results, _ := s.loader.LoadTx(db, selector)

	ctx := model.NewDataContext(results)

	_, ok := model.GetFromContext[*model.Identity](ctx, model.EManager)
	if ok {
		return errors.New("overlapped uuid")
	}

	return nil
}
