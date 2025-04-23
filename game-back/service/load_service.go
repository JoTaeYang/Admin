package service

import (
	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/repo"
)

type LoadService interface {
	Load(id string) error
}

type loadService struct {
	config *config.Configs
	loader *model.Loader
}

func NewLoadService(loader *model.Loader, config *config.Configs) LoadService {
	return &loadService{
		loader: loader,
		config: config,
	}
}

func (s *loadService) Load(id string) error {
	selector := model.NewSelector(id)

	selector.AddSingle(model.EAuth, &repo.ManagerRepository{})

	db := bsql.RDB.GetAdminDB()

	results, err := s.loader.LoadTx(db, selector)
	if err != nil {
		return err
	}

	_ = results

	return nil
}
