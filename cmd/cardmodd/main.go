package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/iamnande/cardmod/internal/config"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/logger"
	"github.com/iamnande/cardmod/internal/repositories"
	"github.com/iamnande/cardmod/internal/server/grpc"
)

var (
	// ServiceName is the name of the service.
	ServiceName = "cardmod"

	// ServiceVersion is the version of the service being deployed.
	// note: this should be overwritten by the linker, using ldflags, during the compilation process.
	ServiceVersion = "v1.0.0-dev"
)

func main() {

	// api: initialize application config
	cfg := config.MustLoad()

	// api: initialize logger
	log := logger.NewLogger(ServiceName, ServiceVersion)

	// api: initialize database client
	dbConn, err := sql.Open("pgx", cfg.DatabaseEndpoint)
	if err != nil {
		log.Error(err, "failed to connect to the database")
		os.Exit(1)
	}
	driver := entsql.OpenDB("postgres", dbConn)
	dbClient := database.NewClient(database.Driver(driver))

	// api: initialize repositories
	cardRepository := repositories.NewCardRepository(dbClient)
	magicRepository := repositories.NewMagicRepository(dbClient)
	calculationRepository := repositories.NewCalculationRepository(dbClient)

	// api: initialize the gRPC server
	grpcServer := grpc.NewServer(&grpc.ServerConfig{
		Logger:  log,
		Port:    cfg.GRPCPort,
		Version: ServiceVersion,

		// Repositories
		CardRepository:        cardRepository,
		MagicRepository:       magicRepository,
		CalculationRepository: calculationRepository,
	})

	// api: start gRPC listener
	go func() {
		log.Info("starting gRPC server")
		if err = grpcServer.Serve(); err != nil {
			log.Error(err, "failed to start gRPC server")
			os.Exit(1)
		}
	}()

	// api: support SIG* signals to handle gracefully stopping gRPC and REST servers
	doneChan := make(chan os.Signal, 1)
	signal.Notify(doneChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	<-doneChan
	log.Info("shutdown signal received")

	// api: setup graceful gRPC server stop
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	go func() {
		log.Info("stopping gRPC server")
		grpcServer.Stop()
		cancel()
	}()

	// api: wait to gracefully stop all the things
	<-ctx.Done()

	// api: we've completely closed everything, goodbye
	log.Info("goodbye, neo.")

}
