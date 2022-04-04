package main

import (
	"github.com/drajk/go-bsb-service/internal"
	"github.com/drajk/go-bsb-service/internal/config"
)

func main() {
	cfg := config.NewApplicationConfig()
	server := internal.SetupServer(cfg)
	server.Start(*cfg)
}
