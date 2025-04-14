package repo

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/model"
)

type ManagerListRepository struct {
}

func (r *ManagerListRepository) Get(tx *sql.Tx) (interface{}, error) {
	queries := []string{
		`SELECT * FROM`,
		bsql.AdminTable,
	}

	resultQuery := strings.Join(queries, " ")

	rows, err := tx.Query(resultQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var m model.Manager
	for rows.Next() {
		if err := rows.Scan(&m.ID, &m.Grade, &m.Name, &m.Password, &m.CreateAt, &m.UpdateAt, &m.Ttl); err != nil {
			return nil, err
		}
		log.Println(m)
	}

	if m.ID == "" {
		return nil, errors.New("not found")
	}

	return &m, nil
}
