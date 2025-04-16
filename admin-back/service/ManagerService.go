package service

import (
	"context"
	"errors"
	"log"

	"github.com/JoTaeYang/Admin/gpkg/bredis"
	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/config"
	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/repo"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ManagerService interface {
	Get(c *gin.Context) ([]*model.Manager, error)
	Put(string, string, string) (interface{}, error)
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

func (s *managerService) Get(c *gin.Context) ([]*model.Manager, error) {
	var results map[model.EModel]interface{}
	var ctx *model.DataContext
	var user []*model.Manager

	checkId, ok := c.Get("userID")
	if !ok {
		return nil, errors.New("not found user id")
	}

	id := checkId.(string)

	selector := model.NewSelector(id)

	repo := &repo.ManagerListRepository{}
	selector.AddRaw(model.EManagerList, repo)

	db := bsql.RDB.GetAdminDB()

	cacheResults, err := s.loader.LoadCache(selector)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(cacheResults) == 0 {
		results, err = s.loader.LoadTx(db, selector)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		ctx = model.NewDataContext(results)

		user, ok = model.GetFromContext[[]*model.Manager](ctx, model.EManagerList)
		if !ok {
			return nil, errors.New("not found datas")
		}

		// cache save
		updateData := []model.IModel{}
		for _, v := range user {
			updateData = append(updateData, v)
		}

		pipe := bredis.R.ReadRedisCli.Pipeline()
		_, err = repo.UpdateCache(model.EManagerList, id, &pipe, updateData)
		if err != nil {
			return nil, err
		}

		_, err = pipe.Exec(context.Background())
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (s *managerService) Put(id, grade, userName string) (interface{}, error) {
	selector := model.NewSelector(id)

	managerRepo := &repo.ManagerRepository{}
	selector.AddSingle(model.EManager, managerRepo)

	db := bsql.RDB.GetAdminDB()

	results, _ := s.loader.LoadTx(db, selector)

	ctx := model.NewDataContext(results)

	user, ok := model.GetFromContext[*model.Manager](ctx, model.EManager)
	if ok {
		if user.ID == id {
			return nil, errors.New("overlapped id")
		}
	}

	// 중복 체크
	defaultPassword := []byte("12345")
	pw, err := bcrypt.GenerateFromPassword(defaultPassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newManager := model.Manager{
		ID:       id,
		Grade:    grade,
		Name:     userName,
		Password: converter.ZeroCopyByteToString(pw),
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	err = managerRepo.Put(tx, &newManager)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return user, nil
}
