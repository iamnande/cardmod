package main

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"os/signal"
	"syscall"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-logr/logr"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/iamnande/cardmod/internal/config"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/logger"
	"github.com/iamnande/cardmod/internal/server"
)

var (
	// ServiceName is the name of the service.
	ServiceName = "cardmodd"

	// ServiceVersion is the version of the service being deployed.
	// note: this should be overwritten by the linker, using ldflags, during the compilation process.
	ServiceVersion = "v0.1.0-alpha"
)

func main() {

	// api: initialize logger
	log := logger.NewLogger(ServiceName, ServiceVersion)

	// api: initialize config
	cfg := config.MustLoad()

	// api: initialize database client
	conn, err := sql.Open("pgx", cfg.Database.DSN())
	if err != nil {
		log.Error(err, "failed to connect to the database")
		os.Exit(-1)
	}
	databaseClient := database.NewClient(database.Driver(entsql.OpenDB("postgres", conn)))

	// api: initialize server
	api := server.NewServer(&server.Config{
		Config:         cfg,
		Logger:         log,
		Version:        ServiceVersion,
		DatabaseClient: databaseClient,
	})

	// api: start the gRPC server
	startServer(cfg, log, api)

	// api: we've completely closed out
	log.Info("goodbye, neo.")

}

// startServer will start the gRPC server and wait for a signal to shutdown.
func startServer(cfg *config.Config, log logr.Logger, api *server.Server) {

	// server: start
	go func() {
		log.Info("starting gRPC server")
		if err := api.Serve(); err != nil {
			log.Error(err, "failed to start gRPC server")
			os.Exit(1)
		}
	}()

	// server: setup shutdown signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)

	// server: watch for shutdown
	<-stopChan
	log.Info("received shutdown signal")

	// server: graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownGracePeriod)
	go func() {
		log.Info("shutting down gRPC server")
		api.GracefullyStop()
		cancel()
	}()
	<-ctx.Done()

	// server: forceful shutdown
	if err := ctx.Err(); err != nil && !errors.Is(err, context.Canceled) {
		api.Stop()
	}

}
