package app

import (
	"github.com/agadilkhan/currency-rate/internal/currency-api/config"
	"github.com/agadilkhan/currency-rate/internal/currency-api/database"
	"log"
)

func Run(cfg *config.Config) {
	db, err := database.NewSQL(cfg.Database.Url)
	if err != nil {
		log.Panicf("cannot connect to db err: %v", err)
	}

	log.Println("database connection success")

	defer func() {
		if err = db.Close(); err != nil {
			log.Panicf("failed to close db err: %v", err)
		}
	}()

}
