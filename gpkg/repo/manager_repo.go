package repo

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/model"
)

type ManagerRepository struct {
}

func (r *ManagerRepository) GetCache(id string) (interface{}, error) {

	return nil, nil
}

func (r *ManagerRepository) Get(tx *sql.Tx, id string) (interface{}, error) {
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
		data.GetTable(),
		`(id, grade, name, password) VALUES (?, ?, ?, ?)`,
	}
	resultQuery := strings.Join(queries, " ")

	_, err := tx.Exec(resultQuery, data.GetCreate()...)
	if err != nil {
		return err
	}

	return nil
}
