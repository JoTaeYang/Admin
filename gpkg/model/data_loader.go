package model

import "database/sql"

type Loader struct {
}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) LoadTx(db *sql.DB, selector *Selector) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	selectList := selector.GetSelect()

	for k, v := range selectList {

	}

	return nil
}
