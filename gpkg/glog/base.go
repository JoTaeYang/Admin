package glog

type BaseLog struct {
	Uid        string `json:"uid"`       // 고유 id
	Srl        string `json:"srl"`       // 로그의 고유 시리얼
	LogType    string `json:"log_type"`  // 로그 타입[META_ITEM, ...]
	LoginSrl   string `json:"login_srl"` // 로그인 세션 srl
	CreateTime string `json:"cre_time"`  // 로그 생성 날짜
}
