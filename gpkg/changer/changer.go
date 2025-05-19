package changer

import (
	"github.com/JoTaeYang/Admin/gpkg/glog"
	"github.com/JoTaeYang/Admin/gpkg/model"
	rf "github.com/JoTaeYang/Admin/gpkg/repo/factory"
)

type Logger = glog.Logger

/*
데이터 타입마다 뭘 만들지는 않고,
뭉칠 수 있는 느낌의 데이터는 뭉쳐서
어떤 새로운 타입을 구현해서
그 쪽에서 데이터 인자를 받아가지구 처리를 한다. 개념임.
*/
type Processor struct {
	*Changer

	Auth     *Auth
	Currecny *Currency
}

type Changer struct {
	*Processor

	hub     *model.ModelHub
	dataCtx *model.DataContext
	updater *model.Updater

	Logger  *Logger
	Factory rf.RepoFactory
	//Table 데이터도 추가
}

func MakeChanger(hub *model.ModelHub, factory rf.RepoFactory, act glog.Act) (*Changer, error) {
	changer := &Changer{
		dataCtx: hub.DataCtx,
		hub:     hub,
		Factory: factory,
	}

	changer.updater = model.NewUpdater()
	changer.Logger = glog.NewLogger(act)
	changer.Processor = &Processor{
		Changer: changer,

		Currecny: newCurrency(changer),
		Auth:     newAuth(changer),
	}

	return changer, nil
}

func (c *Changer) DBExecute() error {
	return c.updater.Execute(c.hub)
}
