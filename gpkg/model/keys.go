package model

import "github.com/JoTaeYang/Admin/gpkg/bsql"

type EModel int32

const (
	EAuth     EModel = 1
	EIdentity EModel = 2
	ECurrency EModel = 3
	EProfile  EModel = 4

	EManager     EModel = 1000
	EManagerList EModel = 1001
)

var EModelMapStr map[EModel]string = map[EModel]string{
	EAuth:     "auth",
	EIdentity: "identity",
	ECurrency: "currency",
	EProfile:  "profile",

	EManager:     "manager",
	EManagerList: "manager_list",
}

var EModelMapInt map[string]EModel = map[string]EModel{
	"auth":     EAuth,
	"identity": EIdentity,
	"currency": ECurrency,
	"profile":  EProfile,

	"manager":      EManager,
	"manager_list": EManagerList,
}

var EModelDBGroup = map[EModel]bsql.DBGroup{
	EAuth:     bsql.DBGroupGame,
	EIdentity: bsql.DBGroupIdentity,
	ECurrency: bsql.DBGroupGame,
	EProfile:  bsql.DBGroupGame,
}
