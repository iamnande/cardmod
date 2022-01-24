package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/database"
	dbCalculation "github.com/iamnande/cardmod/internal/database/calculation"
	"github.com/iamnande/cardmod/internal/database/magic"
	"github.com/iamnande/cardmod/internal/domains/calculation"
)

// calculationRepository is a repository layer for accessing calculation entities in the data layer.
//go:generate mockgen -source ../daos/calculation.go -destination=mocks/calculation.go -package mocks github.com/iamnande/cardmod/internal/daos CalculationDAO
type calculationRepository struct {
	client *database.Client
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.CalculationDAO = (*calculationRepository)(nil)

// NewCalculationRepository initializes a new calculation repository instance.
func NewCalculationRepository(client *database.Client) *calculationRepository {
	return &calculationRepository{
		client: client,
	}
}

// ListCalculations lists all calculation entities. The search in this case will be limited to
// searching against magics, by name.
func (repo *calculationRepository) ListCalculations(ctx context.Context,
	search string) ([]calculation.Calculation, error) {

	// list: search for one specific magic, by name
	result, err := repo.client.Magic.Query().Where(magic.NameEQ(search)).Only(ctx)
	if err != nil {
		return nil, err
	}

	// list: search for calculations that contain magicID
	calculations, err := repo.client.Calculation.Query().Where(dbCalculation.MagicID(result.ID)).All(ctx)
	if err != nil {
		return nil, err
	}

	// list: return collection of calculations
	return marshalCalculationContainers(calculations), nil

}

// CreateCalculation creates a calculation entity.
func (repo *calculationRepository) CreateCalculation(ctx context.Context,
	cardID, magicID uuid.UUID, cardRatio, magicRatio int32) (calculation.Calculation, error) {

	// create: initialize transaction
	tx, err := repo.client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	// create: create the calculation
	calculation, err := tx.Calculation.Create().
		SetCardID(cardID).
		SetMagicID(magicID).
		SetCardRatio(cardRatio).
		SetMagicRatio(magicRatio).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	// create: return the newly created calculation
	return &calculationContainer{calculation}, tx.Commit()

}
