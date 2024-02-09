package main

import (
	"github.com/vpbuyanov/gw-backend-go/internal/server"

	"github.com/vpbuyanov/gw-backend-go/configs"
)

func main() {
	config := configs.LoadConfig()
	runner := server.GetServer(config)

	err := runner.Start()
	if err != nil {
		return
	}
}
