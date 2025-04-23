package repo

import (
	"testing"

	"github.com/JoTaeYang/Admin/gpkg/bsql"
)

func TestRepo(t *testing.T) {
	d := bsql.Config{
		Mode:          "ADMIN",
		Addr:          []string{"127.0.0.1"},
		Port:          []string{"3307"},
		DBName:        "admin",
		Account:       "root",
		Password:      "fncp303151!",
		ShardCount:    1,
		InstanceCount: 1,
	}

	bsql.InitService([]bsql.Config{d})

	r := ManagerRepository{}

	tx, _ := bsql.RDB.GetAdminDB().Begin()
	r.Get(tx, "master_id")
}
