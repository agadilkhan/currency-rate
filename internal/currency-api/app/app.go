package app

import (
	"github.com/agadilkhan/currency-rate/internal/currency-api/config"
	"github.com/agadilkhan/currency-rate/internal/currency-api/controller/http"
	"github.com/agadilkhan/currency-rate/internal/currency-api/database"
	"github.com/agadilkhan/currency-rate/internal/currency-api/repository/postgres"
	"github.com/agadilkhan/currency-rate/internal/currency-api/service"
	"log"
	"os"
	"os/signal"
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

	pgRepo := postgres.New(db.Client)

	srvc := service.New(pgRepo)

	hndlr := http.NewHandler(srvc)

	server := http.NewServer(
		hndlr.InitRouter(),
		http.WithHost(cfg.HttpServer.Host),
		http.WithShutdownTimeout(cfg.HttpServer.ShutdownTimeout),
	)

	server.Start()

	log.Println("http server running...")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

}
