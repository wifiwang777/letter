package main

import (
	"github.com/aisuosuo/letter/api"
	"github.com/aisuosuo/letter/config"
	"github.com/aisuosuo/letter/core/ws"
)

func main() {
	go func() {
		ws.Run()
	}()
	err := api.HttpServer.Run(config.GlobalConfig.Run.HttpAddr)
	if err != nil {
		panic(err)
	}
}
