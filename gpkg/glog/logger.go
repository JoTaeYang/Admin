package glog

type Logger struct {
	Records []string // TODO :: 타입
	Act     Act      // 액션 타입
}

func NewLogger(act Act) *Logger {
	return &Logger{
		Records: make([]string, 0, 5),
		Act:     act,
	}
}

func (l *Logger) InsertMetaItem() {

}
