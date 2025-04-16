package model

import "database/sql"

type Selector struct {
	Id     string
	single map[EModel]ISingleRepository
	raw    map[EModel]IRawRepository
}

type ISingleRepository interface {
	Get(tx *sql.Tx, id string) (interface{}, error)
	GetCache(id string) (interface{}, error)
}

type IRawRepository interface {
	Get(tx *sql.Tx) (interface{}, error)
	GetCache(id string) (interface{}, error)
}

func NewSelector(id string) *Selector {
	return &Selector{
		Id:     id,
		single: make(map[EModel]ISingleRepository, 5),
		raw:    make(map[EModel]IRawRepository, 5),
	}
}

func (s *Selector) AddSingle(key EModel, data ISingleRepository) {
	s.single[key] = data
}

func (s *Selector) AddRaw(key EModel, data IRawRepository) {
	s.raw[key] = data
}

func (s *Selector) GetSingle() map[EModel]ISingleRepository {
	return s.single
}

func (s *Selector) GetRaw() map[EModel]IRawRepository {
	return s.raw
}
