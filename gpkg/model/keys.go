package model

type EModel int32

const (
	EManager     EModel = 1
	EManagerList EModel = 2
)

var EModelMapStr map[EModel]string = map[EModel]string{
	EManager:     "manager",
	EManagerList: "manager_list",
}

var EModelMapInt map[string]EModel = map[string]EModel{
	"manager":      EManager,
	"manager_list": EManagerList,
}
