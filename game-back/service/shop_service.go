package service

import (
	"errors"

	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/changer"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/glog"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/repo"
	"github.com/gin-gonic/gin"
)

type ShopService interface {
	Gacha(c *gin.Context, gachaKey string) (*model.DataContext, error)
}

type shopService struct {
	config *config.Configs
	Loader *model.Loader
}

func (s *shopService) Gacha(c *gin.Context, gachaKey string) (*model.DataContext, error) {
	userID, ok := c.Get("userID")
	if !ok {
		return nil, errors.New("not found user id")
	}

	id := userID.(string)

	identityRepo := repo.IdentityRepository{}

	dbKey := model.EModelDBGroup[model.EIdentity]
	db := bsql.RDB.GetDB(string(dbKey), 0)

	identityInter, err := identityRepo.Get(db, id)
	if err != nil {
		return nil, err
	}

	identity, ok := identityInter.(*model.Identity)
	if !ok {
		return nil, errors.New("not found identity")
	}

	selector := model.NewSelector(id)

	// 소모되는 재화만 read 해서 처리해도 가능.
	// 뽑기니까 어떤 재화를 얻지 않아서 가능하다.
	selector.AddSingle(model.ECurrency, &repo.CurrencyRepository{})

	db = bsql.RDB.GetGameDB(int32(identity.ShardIdx))
	results, err := s.Loader.LoadTx(db, selector)
	if err != nil {
		return nil, err
	}

	dataCtx := model.NewDataContext(results)

	cha, err := changer.MakeChanger(dataCtx, glog.Act_Shop)
	if err != nil {
		return nil, errors.New("changer make error")
	}

	_ = cha

	// 일단 재화 소모 changer 구현하고
	// log format 결정하고
	// log 실어서 나르고
	// log 보기

	return nil, nil
}
