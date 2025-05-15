package model

import (
	"context"
	"database/sql"

	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/gin-gonic/gin"
)

type ModelHub struct {
	Identity *Identity

	ctx    context.Context
	cancel context.CancelFunc

	db       *sql.DB
	selector *Selector

	DataCtx *DataContext
}

func makeIdentity(ctx context.Context, identityRepo ISingleRepository, uid string) *Identity {
	dbKey := EModelDBGroup[EIdentity]

	db := bsql.RDB.GetDB(string(dbKey), 0)

	result, err := identityRepo.Get(ctx, db, uid)
	if err != nil {
		return nil
	}

	identity, ok := result.(*Identity)
	if !ok {
		identity = nil
	}

	return identity
}

func MakeModelHub(c *gin.Context, identityRepo ISingleRepository) *ModelHub {
	userID, ok := c.Get("userID")
	if !ok {
		return nil
	}

	uid := userID.(string)

	ctx, cancel := context.WithCancel(c)

	identity := makeIdentity(ctx, identityRepo, uid)

	//identity save 추가하기

	return &ModelHub{
		Identity: identity,

		ctx:    ctx,
		cancel: cancel,

		db:       bsql.RDB.GetGameDB(int32(identity.ShardIdx)),
		selector: NewSelector(uid),
	}
}

func MakeModelHubAuth(c *gin.Context, id string, identityRepo ISingleRepository) *ModelHub {
	ctx, cancel := context.WithCancel(c)

	//identity save 추가하기
	identity := makeIdentity(ctx, identityRepo, id)

	return &ModelHub{
		Identity: identity,

		ctx:    ctx,
		cancel: cancel,

		db:       bsql.RDB.GetIdentityDB(),
		selector: NewSelector(id),
	}
}

func (m *ModelHub) GetDBCli() *sql.DB {
	return m.db
}
