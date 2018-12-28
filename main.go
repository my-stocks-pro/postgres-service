package main

import (
	"fmt"
	"net/http"
	"os/signal"
	"os"
	"time"
	"context"

	"github.com/my-stocks-pro/postgres-service/infrastructure"
	"github.com/my-stocks-pro/postgres-service/infrastructure/postgres"
	"github.com/my-stocks-pro/postgres-service/engine"
	"github.com/prometheus/common/log"
)

func main() {
	config := infrastructure.NewConfig()

	logger, err := infrastructure.NewLogger()
	if err != nil {
		log.Fatal(err)
		return
	}

	client, err := postgres.NewClient(config)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	persist := postgres.NewPostgres(config, client)

	service := engine.New(config, logger, persist)

	service.InitMux()

	serverHTTP := &http.Server{
		Addr:    config.SPort,
		Handler: service.Engine,
	}

	go func() {
		if err := serverHTTP.ListenAndServe(); err != nil {
			logger.Error(err.Error())
		}
	}()

	signal.Notify(service.QuitOS, os.Interrupt)
	select {
	case <-service.QuitOS:
		logger.Info("Shutdown Postgres server by OS signal...")
	case <-service.QuitRPC:
		logger.Info("Shutdown Postgres server by RPC signal...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := serverHTTP.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Postgres server Shutdown: %s", err.Error()))
	}

	logger.Info("Postgres server exiting")
}
