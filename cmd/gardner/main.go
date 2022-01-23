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
	"github.com/iamnande/cardmod/internal/repositories"
)

var (
	// ServiceName is the name of the service.
	ServiceName = "gardner"

	// ServiceVersion is the version of the service being deployed.
	// note: this should be overwritten by the linker, using ldflags, during the compilation process.
	ServiceVersion = "v1.0.0-dev"
)

func main() {

	// gardner: initialize config
	cfg := config.MustLoad()

	// gardner: initialize logger
	log := logger.NewLogger(ServiceName, ServiceVersion)

	// gardner: initialize database client
	dbConn, err := sql.Open("pgx", cfg.DatabaseEndpoint)
	if err != nil {
		log.Error(err, "failed to connect to the database")
		os.Exit(1)
	}
	driver := entsql.OpenDB("postgres", dbConn)
	dbClient := database.NewClient(database.Driver(driver))

	// gardner: initialize repositories
	ctx := context.TODO()
	cardRepository := repositories.NewCardRepository(dbClient)
	magicRepository := repositories.NewMagicRepository(dbClient)

	// gardner: fetch list of magics to create
	// TODO: handle diff/patching
	magics := LoadMagics()
	log.Info("seeding magics into database")
	for _, new := range magics {

		// check if it exists first
		results, listErr := magicRepository.ListMagics(ctx, new)
		if err != nil {
			log.Error(listErr, "failed to search for %s magic", new)
			os.Exit(1)
		}

		// if it doesn't exist, create it
		if len(results) == 0 {
			_, err = magicRepository.CreateMagic(ctx, new)
			if err != nil {
				log.Error(err, "failed to create %s magic", new)
				os.Exit(1)
			}
		}

	}

	// gardner: fetch list of cards to create
	// TODO: handle diff/patching
	cards := LoadCards()
	log.Info("seeding cards into database")
	for _, new := range cards {

		// check if it exists first
		results, listErr := cardRepository.ListCards(ctx, new)
		if err != nil {
			log.Error(listErr, "failed to search for %s card", new)
			os.Exit(1)
		}

		// if it doesn't exist, create it
		if len(results) == 0 {
			_, err = cardRepository.CreateCard(ctx, new)
			if err != nil {
				log.Error(err, "failed to create %s card")
			}
		}

	}

}
