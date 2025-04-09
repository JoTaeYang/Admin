package model

import "github.com/JoTaeYang/Admin/gpkg/repo"

type Selector struct {
	Id        string
	selection map[string]interface{}
}

func NewSelector(id string) *Selector {
	return &Selector{
		Id:        id,
		selection: make(map[string]interface{}, 5),
	}
}

func AddSelect[T any](selector *Selector, key string, data interface{}) {
	var ok bool
	if _, ok = data.(repo.ISingleRepository[T]); ok {
		selector.addSingle(key, data)
		return
	}
}

func (s *Selector) addSingle(key string, data interface{}) {
	s.selection[key] = data
}

func (s *Selector) GetSelect() map[string]interface{} {
	return s.selection
}
