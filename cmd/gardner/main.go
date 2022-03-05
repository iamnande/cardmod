package main

import (
	"context"
	"database/sql"
	"os"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/iamnande/cardmod/internal/config"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/logger"
)

var (
	// ServiceName is the name of the service.
	ServiceName = "gardner"

	// ServiceVersion is the version of the service being deployed.
	// note: this should be overwritten by the linker, using ldflags, during the compilation process.
	ServiceVersion = "v0.1.0-alpha"
)

func main() {

	// api: initialize logger
	log := logger.NewLogger(ServiceName, ServiceVersion)

	// client: initialize config
	cfg := config.MustLoad()

	// client: initialize connection
	conn, err := sql.Open("pgx", cfg.Database.DSN())
	if err != nil {
		log.Error(err, "failed to connect to the database")
		os.Exit(-1)
	}
	databaseClient := database.NewClient(database.Driver(entsql.OpenDB("postgres", conn)))

	// client: seed cards
	ctx := context.Background()
	if err := seedCards(ctx, databaseClient); err != nil {
		log.Error(err, "failed to seed cards")
		os.Exit(1)
	}

	// client: seed items
	if err := seedItems(ctx, databaseClient); err != nil {
		log.Error(err, "failed to seed items")
		os.Exit(1)
	}

	// client: seed limit breaks
	if err := seedLimitBreaks(ctx, databaseClient); err != nil {
		log.Error(err, "failed to seed limit breaks")
		os.Exit(1)
	}

	// client: seed magic
	if err := seedMagics(ctx, databaseClient); err != nil {
		log.Error(err, "failed to seed magics")
		os.Exit(1)
	}

	// client: seed refinements
	if err := seedRefinements(ctx, databaseClient); err != nil {
		log.Error(err, "failed to seed refinements")
		os.Exit(1)
	}

	// client: all done
	log.Info("seed complete")

}
