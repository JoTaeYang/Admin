package service

import (
	"errors"

	"github.com/JoTaeYang/Admin/gpkg/changer"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/gen"
	"github.com/JoTaeYang/Admin/gpkg/glog"
	"github.com/JoTaeYang/Admin/gpkg/model"
	rf "github.com/JoTaeYang/Admin/gpkg/repo/factory"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	SignUp(c *gin.Context, id string) error
	Login(c *gin.Context, id string) (string, error)
}

type authService struct {
	config  *config.Configs
	loader  *model.Loader
	factory rf.RepoFactory
}

func NewAuthService(loader *model.Loader, config *config.Configs, factory rf.RepoFactory) AuthService {
	return &authService{
		loader:  loader,
		config:  config,
		factory: factory,
	}
}

func (s *authService) SignUp(c *gin.Context, id string) error {
	identityRepo := s.factory.Identity()
	hub := model.MakeModelHubAuth(c, id, &identityRepo)
	if hub == nil {
		return errors.New("Make Hub Error")
	}

	if hub.GetIdentity() != nil {
		return errors.New("overlapped uuid")
	}

	cha, err := changer.MakeChanger(hub, s.factory, glog.Act_Auth)
	if err != nil {
		return errors.New("changer make error")
	}

	cha.Auth.New(id)

	err = cha.DBExecute()
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(c *gin.Context, id string) (string, error) {
	identityRepo := s.factory.Identity()
	hub := model.MakeModelHubAuth(c, id, &identityRepo)
	if hub == nil {
		return "", errors.New("Make Hub Error")
	}

	if hub.GetIdentity() == nil {
		return "", errors.New("not found id")
	}

	return gen.GenerateJWT(id, s.config.GetSecretKey())
}
