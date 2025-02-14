package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Jackabc911/webApi/internal/app/webapi"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/webapi.toml", "path to config file (.toml file)")
}

func main() {
	//Flag parsing and add value to configPath var
	flag.Parse()
	//config instance
	config := webapi.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("can not find path to config, app will use default confs:", err)
	}

	//server instance
	s := webapi.New(config)

	//server start
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
