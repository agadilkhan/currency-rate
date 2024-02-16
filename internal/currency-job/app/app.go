package app

import (
	"context"
	"github.com/agadilkhan/currency-rate/internal/currency-job/config"
	"github.com/agadilkhan/currency-rate/internal/currency-job/database"
	"github.com/agadilkhan/currency-rate/internal/currency-job/repository/postgres"
	"github.com/agadilkhan/currency-rate/internal/currency-job/service"
	"github.com/agadilkhan/currency-rate/internal/currency-job/transport"
	"log"
)

func Run(cfg *config.Config) {
	db, err := database.NewSQL(cfg.Database.Url)
	if err != nil {
		log.Panicf("cannot connect to db")
	}

	log.Println("database connection success")

	pgRepo := postgres.New(db.Client)

	tr := transport.New(cfg.Transport.Host)

	srvc := service.New(pgRepo, *tr)

	srvc.Save(context.TODO())
}
