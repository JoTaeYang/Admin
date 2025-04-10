package repo

type DataContext struct {
	data map[string]interface{}
}

func NewDataContext(raw map[string]interface{}) *DataContext {
	return &DataContext{
		data: raw,
	}
}

func GetFromContext[T any](ctx *DataContext, key string) (T, bool) {
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
