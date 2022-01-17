package main

import (
	"context"
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"

	"github.com/iamnande/cardmod/internal/config"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/repositories"
)

func main() {

	// gardner: initialize config
	cfg := config.MustLoad()

	// gardner: initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// gardner: initialize database client
	dbConn, err := sql.Open("pgx", cfg.DatabaseEndpoint)
	if err != nil {
		logger.Sugar().Fatalf("failed to connect to the database: %v", err)
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
	logger.Info("seeding magics into database")
	for _, new := range magics {

		// check if it exists first
		results, listErr := magicRepository.ListMagics(ctx, new)
		if err != nil {
			logger.Sugar().Fatalf("failed to search for %s magic: %v", new, listErr)
		}

		// if it doesn't exist, create it
		if len(results) == 0 {
			_, err = magicRepository.CreateMagic(ctx, new)
			if err != nil {
				logger.Sugar().Fatalf("failed to create %s magic: %v", new, err)
			}
		}

	}

	// gardner: fetch list of cards to create
	// TODO: handle diff/patching
	cards := LoadCards()
	logger.Info("seeding cards into database")
	for _, new := range cards {

		// check if it exists first
		results, listErr := cardRepository.ListCards(ctx, new)
		if err != nil {
			logger.Sugar().Fatalf("failed to search for %s card: %v", new, listErr)
		}

		// if it doesn't exist, create it
		if len(results) == 0 {
			_, err = cardRepository.CreateCard(ctx, new)
			if err != nil {
				logger.Sugar().Fatalf("failed to create %s card: %v", new, err)
			}
		}

	}

}
