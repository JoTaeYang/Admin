package model

import "database/sql"

type Loader struct {
}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) LoadTx(db *sql.DB, selector *Selector) (map[EModel]interface{}, error) {
	result := make(map[EModel]interface{}, 5)

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	singleList := selector.GetSingle()
	rowList := selector.GetRaw()

	for k, v := range singleList {
		result[k], err = v.Get(tx, selector.Id)
		if err != nil {
			return nil, err
		}
	}

	for k, v := range rowList {
		result[k], err = v.Get(tx)
		if err != nil {
			return nil, err
		}
	}

	tx.Commit()
	return result, nil
}

func (l *Loader) LoadCacheTx(selector *Selector) (map[EModel]interface{}, error) {
	result := make(map[EModel]interface{}, 5)
	var err error

	singleList := selector.GetSingle()
	rowList := selector.GetRaw()

	for k, v := range singleList {
		result[k], err = v.GetCache(selector.Id)
		if err != nil {
			return nil, err
		}
	}

	for k, v := range rowList {
		result[k], err = v.GetCache(selector.Id)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
