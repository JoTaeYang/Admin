package model

import (
	"encoding/json"

	"github.com/JoTaeYang/Admin/gpkg/converter"
)

type IModel interface {
	GetTable() string
	GetCreate() []interface{}
}

func GetJSON(model IModel) string {
	jsonStr, _ := json.Marshal(model)
	return converter.ZeroCopyByteToString(jsonStr)
}
