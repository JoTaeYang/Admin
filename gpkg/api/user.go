package api

import (
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/JoTaeYang/Admin/gpkg/pt"
	"github.com/JoTaeYang/Admin/gpkg/repo"
)

func MakeAccount(id, name, grade string, hub *model.ModelHub) error {
	// 예전에 나는 그냥 한 데이터마다 다 &Data 이렇게 해줬음.

	// rdb에서 pk sk가 필요한가? 지금은 필요가 없는

	// Repository에서 Create를 하려면 데이터가 필요함. 그 데이터를 가져와야 함. 1번.
	// Repository를 Loop 돌면서 Create를 호출해야함.
	// 이것이 기존과 연관이 있는가?

	// Updater가 데이터를 생성해서 집어넣음.
	// 여기선 Repository 들을 인자로 전달해서 처리했음.
	// 그거 그대로 쓰자 그냥
	model.AddSingle(hub, model.EAuth, &repo.AuthRepository{})
	model.AddSingle(hub, model.ECurrency, &repo.CurrencyRepository{})
	model.AddSingle(hub, model.EProfile, &repo.ProfileRepository{})

	updater := model.NewUpdater()

	auth := &model.Auth{
		UserId:   id,
		Grade:    grade,
		ShardIdx: hub.Identity.ShardIdx,
	}

	gold := &model.Currency{
		UserId:       id,
		CurrencyType: int64(pt.Currency_GOLD),
		Count:        0,
	}

	freeCash := &model.Currency{
		UserId:       id,
		CurrencyType: int64(pt.Currency_FREECASH),
		Count:        0,
	}

	cash := &model.Currency{
		UserId:       id,
		CurrencyType: int64(pt.Currency_CASH),
		Count:        0,
	}

	petGacha := &model.Currency{
		UserId:       id,
		CurrencyType: int64(pt.Currency_PET_GACHA_TICKET),
		Count:        0,
	}

	profile := &model.Profile{
		UserId: id,
		Name:   name,
	}

	updater.AddUpsert(auth)
	updater.AddUpsertMulti([]model.IModel{gold, freeCash, cash, petGacha})
	updater.AddUpsert(profile)

	err := updater.Execute(hub)
	if err != nil {
		return err
	}

	return nil
}

func MakeLoadResponse(dataCtx *model.DataContext) []*pt.DataItem {
	return model.ConvertGRPCS(dataCtx)
}
