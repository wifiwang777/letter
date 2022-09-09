package main

import (
	"github.com/aisuosuo/letter/api"
	"github.com/aisuosuo/letter/config"
)

func main() {
	err := api.HttpServer.Run(config.GlobalConfig.Run.HttpAddr)
	if err != nil {
		panic(err)
	}
}
