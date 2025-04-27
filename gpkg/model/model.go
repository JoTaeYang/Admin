package model

import (
	"encoding/json"

	"github.com/JoTaeYang/Admin/gpkg/converter"
)

type IModel interface {
	GetKey() string // DB 테이블 명과 Redis Key에서 씀
	GetEModel() EModel
	GetCreate() []interface{}
}

func GetJSON(model IModel) string {
	jsonStr, _ := json.Marshal(model)
	return converter.ZeroCopyByteToString(jsonStr)
}
