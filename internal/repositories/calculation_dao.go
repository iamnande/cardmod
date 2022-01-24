package repositories

import (
	"github.com/google/uuid"

	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/domains/calculation"
)

// calculationContainer is a wrapper around the data layer schema that satisfies the domain model.
type calculationContainer struct {
	calculation *database.Calculation
}

func (db *calculationContainer) ID() uuid.UUID {
	return db.calculation.ID
}

func (db *calculationContainer) CardID() uuid.UUID {
	return db.calculation.CardID
}

func (db *calculationContainer) MagicID() uuid.UUID {
	return db.calculation.MagicID
}

func (db *calculationContainer) CardRatio() int32 {
	return db.calculation.CardRatio
}

func (db *calculationContainer) MagicRatio() int32 {
	return db.calculation.MagicRatio
}

// marshalCalculationContainers creates containers from the data layer schema.
func marshalCalculationContainers(records []*database.Calculation) []calculation.Calculation {
	calculations := make([]calculation.Calculation, len(records))
	for i := range records {
		calculations[i] = &calculationContainer{records[i]}
	}
	return calculations
}
