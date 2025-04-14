package service

import (
	"log"

	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/repo"
)

type ManagerService interface {
	Get() error
}

type managerService struct {
	config *config.Configs
	loader *model.Loader
}

func NewManagerService(loader *model.Loader, config *config.Configs) ManagerService {
	return &managerService{
		loader: loader,
		config: config,
	}
}

func (s *managerService) Get() error {
	selector := model.NewSelector("")

	selector.AddRaw("manager", &repo.ManagerListRepository{})

	db := bsql.RDB.GetAdminDB()

	results, err := s.loader.LoadTx(db, selector)
	if err != nil {
		log.Println(err)
		return err
	}

	_ = results
	// ctx := repo.NewDataContext(results)

	// user, ok := repo.GetFromContext[*model.Manager](ctx, "manager")
	// if !ok {
	// 	return errors.New("not found datas")
	// }

	return nil
}
