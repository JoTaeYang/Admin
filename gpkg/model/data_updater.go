package model

import (
	"database/sql"
	"errors"
	"log"
)

/*
Selector 에 전달한 key 데이터와
LoadTx로 나온 results 결과를 어떻게 잘 저장한 후에
updater에 들어온 데이터가 results에 있다? => update
없다? => create
이런식으로 하면 어떨까 고민이 된다.
*/

type UpdaterType int

const (
	UpdaterUpsert UpdaterType = iota
	UpdaterDelete
)

type UpdaterEntry struct {
	Type UpdaterType
	Data interface{}
}

type Updater struct {
	updates map[EModel]UpdaterEntry
}

func NewUpdater() *Updater {
	return &Updater{
		updates: make(map[EModel]UpdaterEntry, 5),
	}
}

func (u *Updater) AddUpsert(data IModel) {
	u.updates[data.GetEModel()] = UpdaterEntry{
		Type: UpdaterUpsert,
		Data: data,
	}
}

func (u *Updater) AddUpsertMulti(data []IModel) {
	u.updates[data[0].GetEModel()] = UpdaterEntry{
		Type: UpdaterUpsert,
		Data: data,
	}
}

func (u *Updater) AddDelete(key EModel, data interface{}) {
	u.updates[key] = UpdaterEntry{
		Type: UpdaterDelete,
		Data: data,
	}
}

// todo :: 코드 수정 해보자 뭔가 지저분해
func (u *Updater) Execute(db *sql.DB, selector *Selector) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			log.Printf("rollback failed : %v", err)
		}
	}()

	for key, v := range u.updates {
		entry, err := selector.GetRepository(key)
		if err != nil {
			return errors.New("Get Repository Error")
		}

		switch v.Type {
		case UpdaterUpsert:
			if updater, ok := entry.Repository.(IUpdaterRepository); ok {
				if modelData, ok := v.Data.(IModel); ok {
					err = updater.Update(tx, modelData)
					if err != nil {
						log.Println(err)
						return err
					}
				} else {
					if modelList, ok := v.Data.([]IModel); ok {
						for _, v := range modelList {
							err = updater.Update(tx, v)
							if err != nil {
								log.Println(err)
								return err
							}
						}
					}
				}
			}

		case UpdaterDelete:
			if updater, ok := entry.Repository.(IUpdaterRepository); ok {
				_ = updater
			}
		}
	}

	return tx.Commit()
}
