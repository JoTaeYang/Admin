package model

type EModel int32

const (
	EAuth     EModel = 1
	EIdentity EModel = 2

	EManager     EModel = 1000
	EManagerList EModel = 1001
)

var EModelMapStr map[EModel]string = map[EModel]string{
	EAuth:        "auth",
	EIdentity:    "identity",
	EManager:     "manager",
	EManagerList: "manager_list",
}

var EModelMapInt map[string]EModel = map[string]EModel{
	"auth":         EAuth,
	"identity":     EIdentity,
	"manager":      EManager,
	"manager_list": EManagerList,
}
