package changer

import (
	"github.com/JoTaeYang/Admin/gpkg/glog"
	"github.com/JoTaeYang/Admin/gpkg/model"
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

	Currecny *Currency
}

type Changer struct {
	dataCtx *model.DataContext
	updater *model.Updater

	Logger    *Logger
	Processor *Processor
	//Table 데이터도 추가
}

func MakeChanger(dataCtx *model.DataContext, act glog.Act) (*Changer, error) {
	changer := &Changer{
		dataCtx: dataCtx,
	}

	changer.updater = model.NewUpdater()
	changer.Logger = glog.NewLogger(act)
	changer.Processor = &Processor{
		Changer: changer,

		Currecny: newCurrency(changer),
	}

	return changer, nil
}
