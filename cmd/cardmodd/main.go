package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"

	"github.com/iamnande/cardmod/internal/config"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/repositories"
	"github.com/iamnande/cardmod/internal/server/grpc"
	"github.com/iamnande/cardmod/internal/server/rest"
)

func main() {

	// api: initialize root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// api: initialize application config
	cfg := config.MustLoad()

	// api: initialize logger
	// TODO: test out zapcore and see what value it brings
	// TODO: wrap zap to hide the sugaring, we don't like sugar exposed
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// api: initialize database client
	dbConn, err := sql.Open("pgx", cfg.DatabaseEndpoint)
	if err != nil {
		logger.Sugar().Fatalf("failed to connect to the database: %v", err)
	}
	driver := entsql.OpenDB("postgres", dbConn)
	dbClient := database.NewClient(database.Driver(driver))

	// api: initialize repositories
	cardRepository := repositories.NewCardRepository(dbClient)
	magicRepository := repositories.NewMagicRepository(dbClient)

	// api: initialize the gRPC server
	grpcServer := grpc.NewServer(&grpc.ServerConfig{
		Port:    cfg.GRPCPort,
		Logger:  logger,
		Version: "v1.0.0",

		// Repositories
		CardRepository:  cardRepository,
		MagicRepository: magicRepository,
	})

	// api: initialize REST server
	restServer, err := rest.NewServer(&rest.ServerConfig{
		Context:      ctx,
		Logger:       logger,
		GRPCEndpoint: cfg.GRPCPort,
		RESTEndpoint: cfg.RESTPort,
	})
	if err != nil {
		logger.Sugar().Fatalw("failed to initialize REST server instance", "error", err)
	}

	// api: start gRPC listener
	go func() {
		logger.Info("starting gRPC server")
		if err = grpcServer.Serve(); err != nil {
			logger.Sugar().Fatalw("failed to start gRPC server", "error", err)
		}
	}()

	// api: start REST listener
	go func() {
		logger.Info("starting REST server")
		if err = restServer.Serve(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Sugar().Fatalw("failed to start REST server", "error", err)
			}
		}
	}()

	// api: support SIG* signals to handle gracefully stopping gRPC and REST servers
	doneChan := make(chan os.Signal, 1)
	signal.Notify(doneChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	<-doneChan
	logger.Info("shutdown signal received")
	cancel()

	// api: setup graceful gRPC server stop
	grpcStop := stopServer(func(ctx context.Context) {
		logger.Info("stopping gRPC server")
		grpcServer.Stop()
	})

	// api: setup graceful REST server stop
	restStop := stopServer(func(ctx context.Context) {
		logger.Info("stopping REST server")
		if err = restServer.Stop(ctx); err != nil {
			logger.Sugar().Fatalw("failed to stop REST server", "error", err)
		}
	})

	// api: gracefully stop all the things
	<-grpcStop.Done()
	<-restStop.Done()

	// api: we've completely closed everything, goodbye
	logger.Info("goodbye, neo.")

}

func stopServer(stopHandler func(context.Context)) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	go func() {
		stopHandler(ctx)
		cancel()
	}()
	return ctx
}
