package model

import (
	"github.com/JoTaeYang/Admin/gpkg/pt"
)

type DataContext struct {
	data map[EModel]interface{}
}

type IDataCtx interface {
	ConvertGRPC() *pt.DataItem
}

func NewDataContext(raw map[EModel]interface{}) *DataContext {
	return &DataContext{
		data: raw,
	}
}

func ConvertGRPCS(ctx *DataContext) (lists []*pt.DataItem) {
	lists = make([]*pt.DataItem, 0, 5)
	for _, v := range ctx.data {
		if convert, ok := v.(IDataCtx); ok {
			lists = append(lists, convert.ConvertGRPC())
		} else {
			// 타입 많아지면 switch로 바꾸기
			if tmpDataList, ok := v.([]*Currency); ok {
				for _, tmp := range tmpDataList {
					lists = append(lists, tmp.ConvertGRPC())
				}
			}
		}
	}
	return
}

func GetFromContext[T any](ctx *DataContext, key EModel) (T, bool) {
	val, ok := ctx.data[key]
	if !ok {
		var zero T
		return zero, false
	}

	casted, ok := val.(T)
	if !ok {
		var zero T
		return zero, false
	}

	return casted, true
}

func GetInterfaceFromContext(ctx *DataContext, key EModel) (interface{}, bool) {
	val, ok := ctx.data[key]
	if !ok {
		return nil, false
	}

	return val, true
}
