package main

import (
	"github.com/vpbuyanov/gw-backend-go/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/server"
)

func main() {
	config := configs.LoadConfig()
	runner := server.GetServer(config)

	err := runner.Start()
	if err != nil {
		return
	}
}
