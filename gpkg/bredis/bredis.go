package bredis

import (
	"context"
	"errors"
	"log"

	redis "github.com/redis/go-redis/v9"
)

type Config struct {
	Addr    []string `yaml:"addrs"`
	AppName string   `yaml:"app_name"`
}

type RedisWrap struct {
	WriteRedisCli *redis.ClusterClient
	ReadRedisCli  *redis.ClusterClient
}

var (
	AppName string
	R       RedisWrap
)

func InitService(cfg Config) error {
	AppName = cfg.AppName

	R.WriteRedisCli = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Addr,
		PoolSize: 10,
	})
	if R.WriteRedisCli == nil {
		return errors.New("create redis read cli error")
	}

	R.ReadRedisCli = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Addr,
		PoolSize: 10,
	})
	if R.ReadRedisCli == nil {
		return errors.New("create redis write cli error")
	}
	log.Println(R.ReadRedisCli.Ping(context.Background()))

	return nil
}

func LoadData(key string, argv []string, pipe *redis.Pipeliner) {
	(*pipe).Eval(context.Background(), getSingleData, []string{key}, argv)
}

func LoadZSet(key string, argv []string, pipe *redis.Pipeliner) {
	(*pipe).Eval(context.Background(), getZSetData, []string{key}, argv)
}

func AddZSet(key string, argv []string, pipe *redis.Pipeliner) {
	(*pipe).Eval(context.Background(), setZSetData, []string{key}, argv)
}
