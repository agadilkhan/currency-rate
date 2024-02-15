package main

import (
	"github.com/agadilkhan/currency-rate/internal/currency-api/app"
	"github.com/agadilkhan/currency-rate/internal/currency-api/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config/currency-api")
	if err != nil {
		log.Fatalf("failed to LoadConfig err: %v", err)
	}

	app.Run(cfg)
}
