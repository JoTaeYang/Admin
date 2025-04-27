package bsql

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
Exec() INSERT UPDATE DELETE 에서 사용
결과 값이 필요 없는 SQL 실행에 적합

QueryRow 단일행 조회
반드시 Scan을 통한 값 추출이 필요

Query() 여러 행을 조회
Next로 다음 데이터를 계속 조회
*/

type Config struct {
	Mode          string   `yaml:"mode"`
	Addr          []string `yaml:"addr"`
	Port          []string `yaml:"port"`
	DBName        string   `yaml:"db_name"`
	Account       string   `yaml:"account"`
	Password      string   `yaml:"password"`
	ShardCount    int32    `yaml:"shard_count"`
	InstanceCount int32    `yaml:"instance_count"`
}

type RDBWrap struct {
	dbs map[string][]*sql.DB
}

type DBGroup string

const (
	DBGroupIdentity DBGroup = "IDENTITY"
	DBGroupGame     DBGroup = "GAME"
	DBGroupAdmin    DBGroup = "ADMIN"
)

const (
	DBAuth     string = "AUTH"
	DBIdentity string = "IDENTITY"
)

var (
	cfg           map[string]Config
	RDB           RDBWrap
	AdminTable    string = "manager"
	GameTable     string = "fishtest"
	IdentityTable string = "identity"
)

func InitService(config []Config) error {
	RDB.dbs = make(map[string][]*sql.DB, len(config))
	cfg = make(map[string]Config, len(config))
	for _, v := range config {
		cfg[v.Mode] = v
		for j, _ := range v.Addr {
			dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", v.Account, v.Password, v.Addr[j], v.Port[j], v.DBName)

			db, err := sql.Open("mysql", dataSourceName)
			if err != nil {
				return err
			}

			db.SetMaxOpenConns(50)               // 최대 동시에 열 수 있는 연결 수
			db.SetMaxIdleConns(10)               // idle 상태로 유지할 연결 수
			db.SetConnMaxLifetime(1 * time.Hour) // 연결의 최대 수명

			fmt.Println("mysql db service : ", v.Port[j])

			if RDB.dbs[v.Mode] == nil {
				RDB.dbs[v.Mode] = make([]*sql.DB, len(v.Addr))
			}
			RDB.dbs[v.Mode][j] = db
		}
	}

	return nil
}

func (r *RDBWrap) GetDB(key string, shard int32) *sql.DB {
	return r.getDB(key, shard)
}

func (r *RDBWrap) GetAdminDB() *sql.DB {
	key := "ADMIN"
	return r.getDB(key, cfg[key].ShardCount)
}

func (r *RDBWrap) GetGameDB(shard int32) *sql.DB {
	key := "GAME"
	return r.getDB(key, shard)
}

func (r *RDBWrap) GetIdentityDB() *sql.DB {
	key := "IDENTITY"
	return r.getDB(key, 1)
}

func (r *RDBWrap) GetGameShardIndex() int32 {
	key := "GAME"
	return rand.Int32N(cfg[key].ShardCount) + 1
}

func (r *RDBWrap) getDB(mode string, shard int32) *sql.DB {
	if mode == string(DBGroupIdentity) {
		return r.dbs[mode][0]
	}

	sIdx := int(shard - 1)
	if len(r.dbs[mode]) < sIdx {
		return nil
	}

	return r.dbs[mode][sIdx]
}

func GenerateShardIdx(id int64) int64 {
	key := "GAME"
	log.Println(id, id/10)
	return (id / 10) % int64(cfg[key].ShardCount)
}
