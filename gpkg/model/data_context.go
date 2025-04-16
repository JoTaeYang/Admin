package model

type DataContext struct {
	data map[EModel]interface{}
}

func NewDataContext(raw map[EModel]interface{}) *DataContext {
	return &DataContext{
		data: raw,
	}
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
