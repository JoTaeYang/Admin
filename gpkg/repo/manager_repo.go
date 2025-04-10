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

func (r *ManagerRepository) GetDB(tx *sql.DB, id string) (interface{}, error) {
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
