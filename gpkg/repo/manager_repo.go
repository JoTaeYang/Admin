package repo

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/JoTaeYang/Admin/gpkg/bredis"
	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/redis/go-redis/v9"
)

type ManagerRepository struct {
}

func (r *ManagerRepository) GetCache(key model.EModel, id string, pipe *redis.Pipeliner) (interface{}, error) {
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

func (r *ManagerRepository) Get(db *sql.DB, id string) (interface{}, error) {
	queries := []string{
		`SELECT id, grade, password FROM`,
		bsql.AdminTable,
		`WHERE id = ?`,
	}
	resultQuery := strings.Join(queries, " ")

	rows := db.QueryRow(resultQuery, id)

	var m model.Manager

	if err := rows.Scan(&m.ID, &m.Grade, &m.Password); err != nil {
		return nil, err
	}

	if m.ID == "" {
		return nil, errors.New("not found")
	}

	return &m, nil
}

func (r *ManagerRepository) GetTx(tx *sql.Tx, id string) (interface{}, error) {
	queries := []string{
		`SELECT id, grade, password FROM`,
		bsql.AdminTable,
		`WHERE id = ?`,
	}
	resultQuery := strings.Join(queries, " ")

	rows := tx.QueryRow(resultQuery, id)

	var m model.Manager

	if err := rows.Scan(&m.ID, &m.Grade, &m.Password); err != nil {
		return nil, err
	}

	if m.ID == "" {
		return nil, errors.New("not found")
	}

	return &m, nil
}

func (r *ManagerRepository) Put(tx *sql.Tx, data model.IModel) error {
	queries := []string{
		`INSERT INTO`,
		data.GetKey(),
		`(id, grade, name, password) VALUES (?, ?, ?, ?)`,
	}
	resultQuery := strings.Join(queries, " ")

	_, err := tx.Exec(resultQuery, data.GetCreate()...)
	if err != nil {
		return err
	}

	return nil
}
