package service

import (
	"errors"

	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/gen"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/repo"
)

type AuthService interface {
	SignUp(id string) error
	Login(id string) (string, error)
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
	results, err := s.loader.LoadTx(db, selector)
	if err != nil {
		return err
	}
	ctx := model.NewDataContext(results)

	_, ok := model.GetFromContext[*model.Identity](ctx, model.EIdentity)
	if ok {
		return errors.New("overlapped uuid")
	}

	updater := model.NewUpdater()

	genID := gen.SnowFlake()

	// TODO :: AddUpsert를 Changer에서 해줘도 될 거 같다.
	updater.AddUpsert(&model.Identity{
		ID:       genID,
		UserId:   id,
		ShardIdx: bsql.GenerateShardIdx(genID),
	})

	err = updater.Execute(db, selector)
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) Login(id string) (string, error) {
	selector := model.NewSelector(id)

	selector.AddSingle(model.EIdentity, &repo.IdentityRepository{})

	db := bsql.RDB.GetIdentityDB()
	results, err := s.loader.LoadTx(db, selector)
	if err != nil {
		return "", err
	}
	ctx := model.NewDataContext(results)

	identity, ok := model.GetFromContext[*model.Identity](ctx, model.EIdentity)
	if !ok {
		return "", errors.New("not found id")
	}

	_ = identity

	return gen.GenerateJWT(id, s.config.GetSecretKey())
}
