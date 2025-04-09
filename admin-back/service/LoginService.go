package service

import (
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/repo"
)

type LoginService interface {
	Login(id string, pw string) error
}

type loginService struct {
	loader *model.Loader
}

func NewLoginService(loader *model.Loader) LoginService {
	return &loginService{
		loader: loader,
	}
}

func (s *loginService) Login(id string, pw string) error {
	selector := model.NewSelector(id)

	model.AddSelect(selector, "manager", repo.ManagerRepository{})

	selector.Add("manager")

	err := s.loader.LoadTx(selector)
	if err != nil {
		return err
	}

	return nil
}
