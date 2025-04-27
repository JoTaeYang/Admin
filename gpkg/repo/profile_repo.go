package repo

import (
	"database/sql"
	"strings"

	"github.com/JoTaeYang/Admin/gpkg/bredis"
	"github.com/JoTaeYang/Admin/gpkg/converter"
	"github.com/JoTaeYang/Admin/gpkg/model"
	"github.com/redis/go-redis/v9"
)

type ProfileRepository struct {
}

func (r *ProfileRepository) GetCache(key model.EModel, id string, pipe *redis.Pipeliner) (interface{}, error) {
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

func (r *ProfileRepository) Get(db *sql.DB, id string) (interface{}, error) {
	var m model.Profile
	queries := []string{
		`SELECT user_id, name, name_change_at FROM`,
		m.GetKey(),
		`WHERE user_id = ?`,
	}

	resultQuery := strings.Join(queries, " ")

	rows := db.QueryRow(resultQuery, id)

	if err := rows.Scan(&m.UserId, &m.Name, &m.NameChangeAt); err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *ProfileRepository) GetTx(tx *sql.Tx, id string) (interface{}, error) {
	var m model.Profile
	queries := []string{
		`SELECT user_id, name, name_change_at FROM`,
		m.GetKey(),
		`WHERE user_id = ?`,
	}

	resultQuery := strings.Join(queries, " ")

	rows := tx.QueryRow(resultQuery, id)

	if err := rows.Scan(&m.UserId, &m.Name, &m.NameChangeAt); err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *ProfileRepository) Update(tx *sql.Tx, data model.IModel) error {
	queries := []string{
		`INSERT INTO`,
		data.GetKey(),
		`(user_id, name) VALUES (?, ?)`,
		`ON DUPLICATE KEY UPDATE`,
		`name = VALUES(name),`,
		`name_change_at = IF(name != VALUES(name), CURRENT_TIMESTAMP, name_change_at),`,
		`name_change_count = IF(name != VALUES(name), name_change_count + 1, name_change_count)`,
	}
	resultQuery := strings.Join(queries, " ")

	_, err := tx.Exec(resultQuery, data.GetCreate()...)
	if err != nil {
		return err
	}

	return nil
}
