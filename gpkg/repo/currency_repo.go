package repo

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/JoTaeYang/Admin/gpkg/bredis"
	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/redis/go-redis/v9"
)

type CurrencyRepository struct {
}

func (r *CurrencyRepository) getSQLQuery(table string, option *model.QueryOption) []string {
	queries := []string{
		`SELECT user_id, currency_type, count FROM`,
		table,
		`WHERE user_id = ?`,
	}

	if option != nil && len(option.Params) > 0 {
		// 물음표를 param 개수만큼 생성하고 쉼표로 구분
		placeholders := strings.Repeat("?,", len(option.Params))

		// 마지막 쉼표 제거
		placeholders = placeholders[:len(placeholders)-1]

		queries = append(queries, `AND currency_type IN (`+placeholders+`)`)
	}

	return queries
}

func (r *CurrencyRepository) GetCache(key model.EModel, id string, pipe *redis.Pipeliner) (interface{}, error) {
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

func (r *CurrencyRepository) Get(ctx context.Context, db *sql.DB, id string) (interface{}, error) {
	var m model.Currency

	queries := r.getSQLQuery(m.GetKey(), nil)

	resultQuery := strings.Join(queries, " ")

	rows, err := db.Query(resultQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	mList := make([]*model.Currency, 0, 5)
	for rows.Next() {
		m := &model.Currency{}
		if err := rows.Scan(&m.UserId, &m.CurrencyType, &m.Count); err != nil {
			return nil, err
		}
		mList = append(mList, m)
	}

	return mList, nil
}

func (r *CurrencyRepository) GetTx(tx *sql.Tx, id string) (interface{}, error) {
	var m model.Currency

	queries := r.getSQLQuery(m.GetKey(), nil)

	resultQuery := strings.Join(queries, " ")

	rows, err := tx.Query(resultQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mList := make([]*model.Currency, 0, 5)
	for rows.Next() {
		m := &model.Currency{}
		if err := rows.Scan(&m.UserId, &m.CurrencyType, &m.Count); err != nil {
			return nil, err
		}
		mList = append(mList, m)
	}

	return mList, nil
}

func (r *CurrencyRepository) GetWithOption(tx *sql.Tx, id string, option *model.QueryOption) (interface{}, error) {
	var m model.Currency

	queries := r.getSQLQuery(m.GetKey(), option)

	resultQuery := strings.Join(queries, " ")

	args := []interface{}{id}

	if option != nil && len(option.Params) > 0 {
		for _, v := range option.Params {
			args = append(args, v)
		}
	}

	rows, err := tx.Query(resultQuery, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()
	mList := make([]*model.Currency, 0, 5)
	for rows.Next() {
		m := &model.Currency{}
		if err := rows.Scan(&m.UserId, &m.CurrencyType, &m.Count); err != nil {
			return nil, err
		}
		mList = append(mList, m)
	}

	return mList, nil
}

func (r *CurrencyRepository) Update(tx *sql.Tx, data model.IModel) error {
	queries := []string{
		`INSERT INTO`,
		data.GetKey(),
		`(user_id, currency_type, count) VALUES (?, ?, ?)`,
		`ON DUPLICATE KEY UPDATE count = VALUES(count)`,
	}
	resultQuery := strings.Join(queries, " ")

	_, err := tx.Exec(resultQuery, data.GetCreate()...)
	if err != nil {
		return err
	}

	return nil
}
