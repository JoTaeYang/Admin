package repo

import (
	"database/sql"
	"strings"

	"github.com/JoTaeYang/Admin/gpkg/bredis"
	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/redis/go-redis/v9"
)

type IdentityRepository struct {
}

func (r *IdentityRepository) GetCache(key model.EModel, id string, pipe *redis.Pipeliner) (interface{}, error) {
	id = `{` + id + `}`
	dataKey := model.EModelMapStr[key]
	list := []string{
		bredis.AppName,
		id,
		dataKey,
	}

	keyList := strings.Join(list, ":")

	argv := []string{converter.IntToStr(3600), dataKey}

	bredis.LoadData(keyList, argv, pipe)
	return nil, nil
}

func (r *IdentityRepository) Get(tx *sql.Tx, id string) (interface{}, error) {
	queries := []string{
		`SELECT id, user_id, shard_idx FROM`,
		bsql.IdentityTable,
		`WHERE user_id = ?`,
	}
	resultQuery := strings.Join(queries, " ")

	rows := tx.QueryRow(resultQuery, id)

	var m model.Identity
	if err := rows.Scan(&m.ID, &m.UserId, &m.ShardIdx); err != nil {
		return nil, err
	}

	return &m, nil
}
