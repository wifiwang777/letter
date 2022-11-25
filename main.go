package main

import (
	"github.com/aisuosuo/letter/api"
	"github.com/aisuosuo/letter/config"
	"github.com/aisuosuo/letter/core/ws"
	//_ "github.com/aisuosuo/letter/config/apollo"
)

func main() {
	go func() {
		ws.Run()
	}()
	err := api.HttpServer.Run(config.GlobalConfig.RunConfig.HttpAddr)
	if err != nil {
		panic(err)
	}
}
