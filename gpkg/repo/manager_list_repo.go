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

type ManagerListRepository struct {
}

/*
TODO :: LIMIT와 OFFSET 적용이 필요함
*/
func (r *ManagerListRepository) GetTx(tx *sql.Tx) (interface{}, error) {
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

func (r *ManagerListRepository) Get(db *sql.DB) (interface{}, error) {
	queries := []string{
		`SELECT id, grade, name, created_at, updated_at FROM`,
		bsql.AdminTable,
	}

	resultQuery := strings.Join(queries, " ")

	rows, err := db.Query(resultQuery)
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

func (r *ManagerListRepository) GetCache(key model.EModel, id string, pipe *redis.Pipeliner) (interface{}, error) {
	id = `{` + id + `}`
	dataKey := model.EModelMapStr[key]
	list := []string{
		bredis.AppName,
		id,
		dataKey,
	}

	keyList := strings.Join(list, ":")

	argv := []string{converter.IntToStr(3600), dataKey}

	bredis.LoadZSet(keyList, argv, pipe)
	return nil, nil
}

func (r *ManagerListRepository) UpdateCache(key model.EModel, id string, pipe *redis.Pipeliner, data []model.IModel) (interface{}, error) {
	id = `{` + id + `}`
	dataKey := model.EModelMapStr[key]
	list := []string{
		bredis.AppName,
		id,
		dataKey,
	}

	keyList := strings.Join(list, ":")

	argv := []string{converter.IntToStr(3600)}
	for _, v := range data {
		argv = append(argv, model.GetJSON(v))
	}

	bredis.AddZSet(keyList, argv, pipe)
	return nil, nil
}
