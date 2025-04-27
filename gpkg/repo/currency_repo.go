package repo

import (
	"database/sql"
	"strings"

	"github.com/JoTaeYang/Admin/gpkg/bredis"
	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/redis/go-redis/v9"
)

type CurrencyRepository struct {
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

func (r *CurrencyRepository) Get(db *sql.DB, id string) (interface{}, error) {
	var m model.Currency
	queries := []string{
		`SELECT user_id, currency_type, count FROM`,
		m.GetKey(),
		`WHERE user_id = ?`,
	}

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
	queries := []string{
		`SELECT user_id, currency_type, count FROM`,
		m.GetKey(),
		`WHERE user_id = ?`,
	}

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
