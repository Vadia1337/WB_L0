package main

import (
	"WB_L0/internal/app/server"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	ConfigPath string = "configs/serverconf.toml"
)

func main() {

	config := &server.Config{}
	_, err := toml.DecodeFile(ConfigPath, config)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(config)

	srv.Start()
}
