package daos

import (
	"context"

	"github.com/google/uuid"

	"github.com/iamnande/cardmod/internal/domains/calculation"
)

// CalculationDAO is the data access object (controller<->data layer) interface definition.
type CalculationDAO interface {

	// ListCalculations lists all calculation entities.
	ListCalculations(ctx context.Context, search string) ([]calculation.Calculation, error)

	// CreateCalculation creates a calculation entity.
	CreateCalculation(ctx context.Context,
		cardID, magicID uuid.UUID, cardRatio, magicRatio int32) (calculation.Calculation, error)
}
