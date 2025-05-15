package service

import (
	"errors"

	"github.com/JoTaeYang/Admin/gpkg/changer"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/glog"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/pt"
	"github.com/JoTaeYang/Admin/gpkg/repo"
	"github.com/gin-gonic/gin"
)

type ShopService interface {
	Gacha(c *gin.Context, gachaKey string, count int32) (*model.DataContext, error)
}

type shopService struct {
	config *config.Configs
	Loader *model.Loader
}

func NewShopService(Loader *model.Loader, config *config.Configs) UserService {
	return &userService{
		Loader: Loader,
		config: config,
	}
}

func (s *shopService) Gacha(c *gin.Context, gachaKey string, count int32) (*model.DataContext, error) {
	hub := model.MakeModelHub(c, &repo.IdentityRepository{})
	if hub == nil {
		return nil, errors.New("Make Hub Error")
	}

	// 소모되는 재화만 read 해서 처리해도 가능.
	// 뽑기니까 어떤 재화를 얻지 않아서 가능하다.
	model.AddSingle(hub, model.ECurrency, &repo.CurrencyRepository{})

	err := s.Loader.LoadTx(hub)
	if err != nil {
		return nil, err
	}

	cha, err := changer.MakeChanger(hub, glog.Act_Shop)
	if err != nil {
		return nil, errors.New("changer make error")
	}

	cha.Processor.Currecny.Use(pt.Currency_PET_GACHA_TICKET, int64(count))

	return nil, nil
}
