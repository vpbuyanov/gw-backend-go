package main

import (
	"fmt"

	"github.com/vpbuyanov/gw-backend-go/configs"
)

func main() {
	config := configs.LoadConfig()

	fmt.Println(config)
}
