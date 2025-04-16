package model

type EModel int32

const (
	EManager EModel = 1
)

var EModelMapStr map[EModel]string = map[EModel]string{
	EManager: "manager",
}

var EModelMapInt map[string]EModel = map[string]EModel{
	"manager": EManager,
}
