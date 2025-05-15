package repo

import (
	"context"
	"database/sql"
	"strings"

	"github.com/JoTaeYang/Admin/gpkg/bredis"
	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/redis/go-redis/v9"
)

type AuthRepository struct {
}

func (r *AuthRepository) GetCache(key model.EModel, id string, pipe *redis.Pipeliner) (interface{}, error) {
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

func (r *AuthRepository) Get(ctx context.Context, db *sql.DB, id string) (interface{}, error) {
	var m model.Auth
	queries := []string{
		`SELECT id, user_id FROM`,
		m.GetKey(),
		`WHERE user_id = ?`,
	}
	resultQuery := strings.Join(queries, " ")

	rows := db.QueryRowContext(ctx, resultQuery, id)
	if ctx.Err() != nil {
		return m, ctx.Err()
	}

	if err := rows.Scan(&m.ID, &m.UserId); err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *AuthRepository) GetTx(ctx context.Context, tx *sql.Tx, id string) (interface{}, error) {
	var m model.Identity
	queries := []string{
		`SELECT id, user_id, shard_idx FROM`,
		m.GetKey(),
		`WHERE user_id = ?`,
	}
	resultQuery := strings.Join(queries, " ")

	rows := tx.QueryRowContext(ctx, resultQuery, id)

	if err := rows.Scan(&m.ID, &m.UserId, &m.ShardIdx); err != nil {
		return nil, err
	}

	if ctx.Err() != nil {
		return m, ctx.Err()
	}

	return &m, nil
}

func (r *AuthRepository) Update(ctx context.Context, tx *sql.Tx, data model.IModel) error {
	queries := []string{
		`INSERT INTO`,
		data.GetKey(),
		`(user_id, grade, shard_idx) VALUES (?, ?, ?)`,
	}
	resultQuery := strings.Join(queries, " ")

	_, err := tx.ExecContext(ctx, resultQuery, data.GetCreate()...)
	if err != nil {
		return err
	}

	return nil
}
