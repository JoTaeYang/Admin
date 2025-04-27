package repo

import (
	"database/sql"
	"strings"

	"github.com/JoTaeYang/Admin/gpkg/bredis"
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

func (r *IdentityRepository) GetTx(tx *sql.Tx, id string) (interface{}, error) {
	var m model.Identity
	queries := []string{
		`SELECT id, user_id, shard_idx FROM`,
		m.GetKey(),
		`WHERE user_id = ?`,
	}
	resultQuery := strings.Join(queries, " ")

	rows := tx.QueryRow(resultQuery, id)

	if err := rows.Scan(&m.ID, &m.UserId, &m.ShardIdx); err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *IdentityRepository) Get(db *sql.DB, id string) (interface{}, error) {
	var m model.Identity
	queries := []string{
		`SELECT id, user_id, shard_idx FROM`,
		m.GetKey(),
		`WHERE user_id = ?`,
	}
	resultQuery := strings.Join(queries, " ")

	rows := db.QueryRow(resultQuery, id)

	if err := rows.Scan(&m.ID, &m.UserId, &m.ShardIdx); err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *IdentityRepository) Update(tx *sql.Tx, data model.IModel) error {
	queries := []string{
		`INSERT INTO`,
		data.GetKey(),
		`(id, user_id, shard_idx) VALUES (?, ?, ?)`,
	}
	resultQuery := strings.Join(queries, " ")

	_, err := tx.Exec(resultQuery, data.GetCreate()...)
	if err != nil {
		return err
	}

	return nil
}
