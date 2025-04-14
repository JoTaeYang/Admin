package model

import "database/sql"

type Selector struct {
	Id     string
	single map[string]ISingleRepository
	raw    map[string]IRawRepository
}

type ISingleRepository interface {
	Get(tx *sql.Tx, id string) (interface{}, error)
}

type IRawRepository interface {
	Get(tx *sql.Tx) (interface{}, error)
}

func NewSelector(id string) *Selector {
	return &Selector{
		Id:     id,
		single: make(map[string]ISingleRepository, 5),
	}
}

func (s *Selector) AddSingle(key string, data ISingleRepository) {
	s.single[key] = data
}

func (s *Selector) AddRaw(key string, data IRawRepository) {
	s.raw[key] = data
}

func (s *Selector) GetSingle() map[string]ISingleRepository {
	return s.single
}

func (s *Selector) GetRaw() map[string]IRawRepository {
	return s.raw
}
