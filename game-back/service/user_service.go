package service

import (
	"errors"
	"log"

	"github.com/JoTaeYang/Admin/gpkg/api"
	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/repo"
)

type UserService interface {
	Load(id string) (*model.DataContext, error)
	New(id, name string) error
}

type userService struct {
	config *config.Configs
	Loader *model.Loader
}

func NewUserService(Userer *model.Loader, config *config.Configs) UserService {
	return &userService{
		Loader: Userer,
		config: config,
	}
}

func (s *userService) Load(id string) (*model.DataContext, error) {
	selector := model.NewSelector(id)

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

	selector.AddSingle(model.EAuth, &repo.AuthRepository{})

	selector.AddSingle(model.ECurrency, &repo.CurrencyRepository{})
	//selector.AddSingleWithOption(model.ECurrency, &repo.CurrencyRepository{}, &model.QueryOption{Params: []string{"1", "2"}})
	selector.AddSingle(model.EProfile, &repo.ProfileRepository{})

	db = bsql.RDB.GetGameDB(int32(identity.ShardIdx))
	results, err := s.Loader.LoadTx(db, selector)
	if err != nil {
		return nil, err
	}

	// identity save 하기

	return model.NewDataContext(results), nil
}

func (s *userService) New(id, name string) error {
	selector := model.NewSelector(id)

	identityRepo := repo.IdentityRepository{}

	dbKey := model.EModelDBGroup[model.EIdentity]
	db := bsql.RDB.GetDB(string(dbKey), 0)
	identityInter, err := identityRepo.Get(db, id)
	if err != nil {
		return err
	}

	identity, ok := identityInter.(*model.Identity)
	if !ok {
		return errors.New("not found identity")
	}

	selector.AddSingle(model.EAuth, &repo.AuthRepository{})
	results, err := s.Loader.Load(selector, identity.ShardIdx)
	if err != nil {
		return err
	}

	dataCtx := model.NewDataContext(results)

	_, ok = model.GetFromContext[*model.Auth](dataCtx, model.EAuth)
	if ok {
		return errors.New("overlapped auth")
	}

	log.Println("shard count : ", identity.ShardIdx)
	err = api.MakeAccount(id, name, "user", identity.ShardIdx, selector)
	if err != nil {
		return err
	}

	return nil
}
