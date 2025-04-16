package repo

import (
	"database/sql"
	"strings"

	"github.com/JoTaeYang/Admin/gpkg/bsql"
	"github.com/JoTaeYang/Admin/gpkg/model"
)

type ManagerListRepository struct {
}

func (r *ManagerListRepository) Get(tx *sql.Tx) (interface{}, error) {
	queries := []string{
		`SELECT id, grade, name, created_at, updated_at FROM`,
		bsql.AdminTable,
	}

	resultQuery := strings.Join(queries, " ")

	rows, err := tx.Query(resultQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mList := make([]*model.Manager, 0, 5)
	for rows.Next() {
		m := model.Manager{}
		if err := rows.Scan(&m.ID, &m.Grade, &m.Name, &m.CreateAt, &m.UpdateAt); err != nil {
			return nil, err
		}
		mList = append(mList, &m)
	}

	return mList, nil
}

func (r *ManagerListRepository) GetCache(id string) (interface{}, error) {
	return nil, nil
}
