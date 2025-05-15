package model

import (
	"context"
	"database/sql"
	"errors"

	"github.com/redis/go-redis/v9"
)

type SelectionType int

const (
	SelectionTypeSingle SelectionType = iota
	SelectionTypeMulti
)

type QueryOption struct {
	Params []string
}

type SelectionEntry struct {
	Type       SelectionType
	Repository interface{}
	Option     *QueryOption
}

type Selector struct {
	Id         string
	selections map[EModel]SelectionEntry
}

type ISingleRepository interface {
	GetTx(ctx context.Context, tx *sql.Tx, id string) (interface{}, error)
	Get(ctx context.Context, db *sql.DB, id string) (interface{}, error)
	GetCache(key EModel, id string, pipe *redis.Pipeliner) (interface{}, error)
}

type IOptionRepository interface {
	GetWithOption(ctx context.Context, tx *sql.Tx, id string, option *QueryOption) (interface{}, error)
}

type IMultiRepository interface {
	GetTx(ctx context.Context, tx *sql.Tx) (interface{}, error)
	Get(c context.Context, db *sql.DB) (interface{}, error)
	GetCache(key EModel, id string, pipe *redis.Pipeliner) (interface{}, error) // THINK : id를 slice로 받아야 할까? 그런 경우가 있을까.. 데이터가 없으니 너무 턱 막히네
}

type IUpdaterRepository interface {
	Update(ctx context.Context, tx *sql.Tx, data IModel) error
}

func NewSelector(id string) *Selector {
	return &Selector{
		Id:         id,
		selections: make(map[EModel]SelectionEntry, 5),
	}
}

func (s *Selector) AddSingleWithOption(key EModel, data ISingleRepository, option *QueryOption) {
	if _, ok := s.selections[key]; ok {
		return
	}

	s.selections[key] = SelectionEntry{
		Type:       SelectionTypeSingle,
		Repository: data,
		Option:     option,
	}
}

func AddSingle(hub *ModelHub, key EModel, data ISingleRepository) {
	s := hub.selector
	if _, ok := s.selections[key]; ok {
		return
	}

	s.selections[key] = SelectionEntry{
		Type:       SelectionTypeSingle,
		Repository: data,
	}
}

func (s *Selector) AddSingle(key EModel, data ISingleRepository) {
	if _, ok := s.selections[key]; ok {
		return
	}

	s.selections[key] = SelectionEntry{
		Type:       SelectionTypeSingle,
		Repository: data,
	}
}

func (s *Selector) AddMulti(key EModel, data IMultiRepository) {
	s.selections[key] = SelectionEntry{
		Type:       SelectionTypeMulti,
		Repository: data,
	}
}

func (s *Selector) GetSelect() map[EModel]SelectionEntry {
	return s.selections
}

func (s *Selector) GetRepository(key EModel) (*SelectionEntry, error) {
	var data SelectionEntry
	var ok bool
	if data, ok = s.selections[key]; !ok {
		return nil, errors.New("not found")
	}
	return &data, nil
}
