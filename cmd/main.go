package main

import (
	"github.com/agadilkhan/currency-rate/internal/app"
	"github.com/agadilkhan/currency-rate/internal/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config")
	if err != nil {
		log.Fatalf("failed to LoadConfig err: %v", err)
	}

	app.Run(cfg)
}
