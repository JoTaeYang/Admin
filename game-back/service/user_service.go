package service

import (
	"errors"

	"github.com/JoTaeYang/Admin/gpkg/api"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/model"
	rf "github.com/JoTaeYang/Admin/gpkg/repo/factory"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Load(c *gin.Context, id string) (*model.DataContext, error)
	New(c *gin.Context, id, name string) error
}

type userService struct {
	config  *config.Configs
	Loader  *model.Loader
	factory rf.RepoFactory
}

func NewUserService(Loader *model.Loader, config *config.Configs, factory rf.RepoFactory) UserService {
	return &userService{
		Loader:  Loader,
		config:  config,
		factory: factory,
	}
}

func (s *userService) Load(c *gin.Context, id string) (*model.DataContext, error) {
	identityRepo := s.factory.Identity()
	authRepo := s.factory.Auth()
	currencyRepo := s.factory.Currency()
	profileRepo := s.factory.Profile()

	hub := model.MakeModelHub(c, &identityRepo)
	if hub == nil {
		return nil, errors.New("Make Hub Error")
	}

	model.AddSingle(hub, model.EAuth, &authRepo)
	model.AddSingle(hub, model.ECurrency, &currencyRepo)
	model.AddSingle(hub, model.EProfile, &profileRepo)

	//selector.AddSingleWithOption(model.ECurrency, &repo.CurrencyRepository{}, &model.QueryOption{Params: []string{"1", "2"}})

	err := s.Loader.LoadTx(hub)
	if err != nil {
		return nil, err
	}

	return model.NewDataContextHub(hub), nil
}

func (s *userService) New(c *gin.Context, id, name string) error {
	identityRepo := s.factory.Identity()
	authRepo := s.factory.Auth()
	hub := model.MakeModelHub(c, &identityRepo)
	if hub == nil {
		return errors.New("Make Hub Error")
	}

	model.AddSingle(hub, model.EAuth, &authRepo)
	err := s.Loader.Load(hub)
	if err != nil {
		return err
	}

	_, ok := hub.GetAuth()
	if ok {
		return errors.New("overlapped auth")
	}

	err = api.MakeAccount(id, name, "user", hub, s.factory)
	if err != nil {
		return err
	}
	return nil
}
