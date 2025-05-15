package service

import (
	"errors"

	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/gen"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/repo"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	SignUp(c *gin.Context, id string) error
	Login(c *gin.Context, id string) (string, error)
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

func (s *authService) SignUp(c *gin.Context, id string) error {
	hub := model.MakeModelHubAuth(c, &repo.IdentityRepository{})
	if hub == nil {
		return errors.New("Make Hub Error")
	}

	model.AddSingle(hub, model.EIdentity, &repo.IdentityRepository{})

	err := s.loader.LoadTx(hub)
	if err != nil {
		return err
	}

	_, ok := hub.GetIdentity()
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

	err = updater.Execute(hub)
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) Login(c *gin.Context, id string) (string, error) {
	hub := model.MakeModelHubAuth(c, &repo.IdentityRepository{})
	if hub == nil {
		return "", errors.New("Make Hub Error")
	}

	model.AddSingle(hub, model.EIdentity, &repo.IdentityRepository{})

	err := s.loader.LoadTx(hub)
	if err != nil {
		return "", err
	}

	_, ok := hub.GetIdentity()
	if !ok {
		return "", errors.New("not found id")
	}

	return gen.GenerateJWT(id, s.config.GetSecretKey())
}
