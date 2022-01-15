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
	ctx := context.Background()
	cardRepository := repositories.NewCardRepository(dbClient)
	magicRepository := repositories.NewMagicRepository(dbClient)
	calculationRepository := repositories.NewCalculationRepository(dbClient)

	// gardner: seed the database with cards, magics, and calculations (card<->magic ratios)
	log.Info("seeding data into database")
	for i := 0; i < len(calculations); i++ {

		// ratio to operate on (for simplicity)
		ratio := calculations[i]

		// create card record
		card, err := cardRepository.CreateCard(ctx, ratio.card)
		if err != nil {
			log.Error(err, "failed to create card")
			os.Exit(1)
		}

		// create magic record
		magic, err := magicRepository.CreateMagic(ctx, ratio.magic)
		if err != nil {
			log.Error(err, "failed to create magic")
			os.Exit(1)
		}

		// create the calcuation record (card<->magic ratio)
		_, err = calculationRepository.CreateCalculation(ctx, card.ID(), magic.ID(), ratio.cardRatio, ratio.magicRatio)
		if err != nil {
			log.Error(err, "failed to create calculation")
			os.Exit(1)
		}

	}

}
