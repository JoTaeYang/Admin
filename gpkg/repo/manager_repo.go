package repo

import (
	"database/sql"

	"github.com/JoTaeYang/Admin/gpkg/model"
)

type ManagerRepository struct {
}

func (r *ManagerRepository) Get(tx *sql.Tx, id string) (*model.Manager, error) {

	return nil, nil
}
