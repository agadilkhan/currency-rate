package app

import (
	"context"
	"github.com/agadilkhan/currency-rate/internal/config"
	"github.com/agadilkhan/currency-rate/internal/controller/http"
	"github.com/agadilkhan/currency-rate/internal/job"
	"github.com/agadilkhan/currency-rate/internal/repository/postgres"
	"github.com/agadilkhan/currency-rate/internal/service"
	"github.com/agadilkhan/currency-rate/internal/transport"
	"github.com/agadilkhan/currency-rate/pkg/database"
	"github.com/agadilkhan/currency-rate/pkg/httpserver"
	"log"
	"os"
	"os/signal"
)

// Run initialize whole application.
func Run(cfg *config.Config) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

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

	tr := transport.New(cfg.Host)

	srvc := service.New(pgRepo, tr)

	hndlr := http.NewHandler(srvc)

	server := httpserver.New(
		hndlr.InitRouter(),
		httpserver.WithHost(cfg.HttpServer.Port),
		httpserver.WithShutdownTimeout(cfg.HttpServer.ShutdownTimeout),
	)

	server.Start()

	log.Println("http server running...")

	jb := job.New(srvc, cfg.Job.UpdateInterval)
	go jb.Run(ctx)

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
