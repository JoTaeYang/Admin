package main

import (
	"log"

	"github.com/JoTaeYang/Admin/gpkg/config"
)

var (
	configDefaultPath = "./"
	configDefaultName = "config.yaml"
	cfg               = config.Configs{}
)

func InitConfig() (err error) {
	err = config.InitConfig(&cfg, configDefaultPath+configDefaultName)
	if err != nil {
		return
	}

	log.Println(cfg)
	return nil
}
