package model

import (
	"context"
	"database/sql"
	"log"

	"github.com/JoTaeYang/Admin/gpkg/bredis"
	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/redis/go-redis/v9"
)

type Loader struct {
}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) Load(hub *ModelHub) error {

	result := make(map[EModel]interface{}, 5)

	selector := hub.selector

	selections := selector.GetSelect()
	for key, entry := range selections {
		dbKey := EModelDBGroup[key]

		db := bsql.RDB.GetDB(string(dbKey), int32(hub.Identity.ShardIdx))
		switch entry.Type {
		case SelectionTypeSingle:
			repo := entry.Repository.(ISingleRepository)
			result[key], _ = repo.Get(hub.ctx, db, selector.Id)
			// case SelectionTypeRange:
			// 	repo := entry.Repository.(IMultiRepository)
			// 	result[key], _ = repo.Get(hub.ctx, db)
		}
	}

	MakeDataContext(hub, result)

	return nil
}

func (l *Loader) LoadTx(hub *ModelHub) error {
	result := make(map[EModel]interface{}, 5)

	tx, err := hub.db.BeginTx(hub.ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			log.Printf("rollback failed : %v", err)
		}
	}()

	selector := hub.selector
	selections := selector.GetSelect()
	for key, entry := range selections {
		switch entry.Type {
		case SelectionTypeSingle:
			if entry.Option == nil {
				repo := entry.Repository.(ISingleRepository)
				result[key], err = repo.GetTx(hub.ctx, tx, selector.Id)
				if err != nil {
					return err
				}
			} else {
				repo := entry.Repository.(IOptionRepository)
				result[key], _ = repo.GetWithOption(hub.ctx, tx, selector.Id, entry.Option)
				if err != nil {
					return err
				}
			}
		case SelectionTypeRange:
			// repo := entry.Repository.(IMultiRepository)
			// result[key], err = repo.GetTx(hub.ctx, tx)
			// if err != nil {
			// 	return err
			// }
		}
	}

	tx.Commit()

	MakeDataContext(hub, result)

	return nil
}

func (l *Loader) LoadCache(selector *Selector) (map[EModel]interface{}, error) {
	result := make(map[EModel]interface{}, 5)
	var err error

	pipe := bredis.R.ReadRedisCli.Pipeline()

	selections := selector.GetSelect()

	for key, entry := range selections {
		switch entry.Type {
		case SelectionTypeSingle:
			repo := entry.Repository.(ISingleRepository)
			result[key], err = repo.GetCache(key, selector.Id, &pipe)
			if err != nil {
				return nil, err
			}
		case SelectionTypeRange:
			// repo := entry.Repository.(IMultiRepository)
			// result[key], err = repo.GetCache(key, selector.Id, &pipe)
			// if err != nil {
			// 	return nil, err
			// }
		}
	}

	results, err := pipe.Exec(context.Background())
	if err != nil {
		return nil, err
	}

	for _, v := range results {
		_, ok := v.(*redis.Cmd)
		if !ok {
			continue
		}

		redisResult, err := v.(*redis.Cmd).Result()
		if err != nil {
			continue
		}

		r, ok := redisResult.([]interface{})
		if ok {
			resultDatas, ok := r[0].([]interface{})
			if ok {
				if len(resultDatas) < 2 {
					continue
				}
				keyName := resultDatas[0].(string)

				log.Println(EModelMapInt[keyName])
			}

		}
	}

	return make(map[EModel]interface{}), nil
}
