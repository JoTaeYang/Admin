package model

import "database/sql"

type Selector struct {
	Id        string
	selection map[string]ISingleRepository
}

type ISingleRepository interface {
	Get(tx *sql.Tx, id string) (interface{}, error)
}

func NewSelector(id string) *Selector {
	return &Selector{
		Id:        id,
		selection: make(map[string]ISingleRepository, 5),
	}
}

func (s *Selector) AddSingle(key string, data ISingleRepository) {
	s.selection[key] = data
}

func (s *Selector) GetSelect() map[string]ISingleRepository {
	return s.selection
}
