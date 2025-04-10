package model

import "database/sql"

type Loader struct {
}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) LoadTx(db *sql.DB, selector *Selector) (map[string]interface{}, error) {
	result := make(map[string]interface{}, 5)

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	selectList := selector.GetSelect()

	for k, v := range selectList {
		result[k], err = v.Get(tx, selector.Id)
		if err != nil {
			return nil, err
		}
	}

	tx.Commit()
	return result, nil
}
