package model

import (
	"context"
	"database/sql"
	"log"

	"github.com/JoTaeYang/Admin/gpkg/bredis"
	"github.com/redis/go-redis/v9"
)

type Loader struct {
}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) LoadTx(db *sql.DB, selector *Selector) (map[EModel]interface{}, error) {
	result := make(map[EModel]interface{}, 5)

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	singleList := selector.GetSingle()
	rowList := selector.GetRaw()

	for k, v := range singleList {
		result[k], err = v.Get(tx, selector.Id)
		if err != nil {
			return nil, err
		}
	}

	for k, v := range rowList {
		result[k], err = v.Get(tx)
		if err != nil {
			return nil, err
		}
	}

	tx.Commit()
	return result, nil
}

func (l *Loader) LoadCache(selector *Selector) (map[EModel]interface{}, error) {
	result := make(map[EModel]interface{}, 5)
	var err error

	pipe := bredis.R.ReadRedisCli.Pipeline()

	singleList := selector.GetSingle()
	rowList := selector.GetRaw()

	for k, v := range singleList {
		result[k], err = v.GetCache(k, selector.Id, &pipe)
		if err != nil {
			return nil, err
		}
	}

	for k, v := range rowList {
		result[k], err = v.GetCache(k, selector.Id, &pipe)
		if err != nil {
			return nil, err
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
