package glog

type Act int32

const (
	Act_Shop = 1000
)

type MetaItem struct {
	BaseLog
	ActType      int32  `json:"act_type"`  // 변동 사유
	OperatorType int8   `json:"oper_type"` // 1 : add / -1 sub
	Amt          int64  `json:"amt"`       // 변화된 수량
	AftAmt       int64  `json:"aft_amt"`   // 변동 후 수량
	ItemType     int32  `json:"item_type"` // 아이템 종류류
	ItemId       string `json:"item_id"`   // 아이템 Id
}
