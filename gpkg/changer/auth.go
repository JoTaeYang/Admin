package changer

import (
	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/gen"
	"github.com/JoTaeYang/Admin/gpkg/model"
)

type Auth struct {
	*Changer
}

func newAuth(changer *Changer) *Auth {
	return &Auth{
		Changer: changer,
	}
}

func (a *Auth) New(id string) {
	genID := gen.SnowFlake()

	identityRepo := a.Factory.Identity()
	a.updater.AddUpsert(&model.Identity{
		ID:       genID,
		UserId:   id,
		ShardIdx: bsql.GenerateShardIdx(genID),
	}, &identityRepo)
}
