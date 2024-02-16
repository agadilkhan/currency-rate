package main

import (
	"github.com/agadilkhan/currency-rate/internal/currency-job/app"
	"github.com/agadilkhan/currency-rate/internal/currency-job/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config/currency-job")
	if err != nil {
		log.Fatalf("failed to LoadConfig err: %v", err)
	}

	app.Run(cfg)
}
