package changer

import "github.com/JoTaeYang/Admin/gpkg/glog"

type Logger = glog.Logger

/*
데이터 타입마다 뭘 만들지는 않고,
뭉칠 수 있는 느낌의 데이터는 뭉쳐서
어떤 새로운 타입을 구현해서
그 쪽에서 데이터 인자를 받아가지구 처리를 한다. 개념임.
*/
type Changer struct {
	Logger *Logger
}
