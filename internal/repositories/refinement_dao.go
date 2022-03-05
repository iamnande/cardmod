package repositories

import (
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/models"
)

// refinementDAOWrapper is the concrete instance of a refinement.
type refinementDAOWrapper struct {
	refinement *database.Refinement
}

func (db *refinementDAOWrapper) Source() string {
	return db.refinement.Source
}

func (db *refinementDAOWrapper) Target() string {
	return db.refinement.Target
}

func (db *refinementDAOWrapper) Numerator() int32 {
	return db.refinement.Numerator
}

func (db *refinementDAOWrapper) Denominator() int32 {
	return db.refinement.Denominator
}

// marshalRefinementDAO creates a DAO instance from the database schema.
func marshalRefinementDAO(record *database.Refinement) models.Refinement {
	return &refinementDAOWrapper{record}
}

// marshalRefinementDAO creates a DAO instance from the database schema.
func marshalRefinementDAOs(records []*database.Refinement) []models.Refinement {
	refinements := make([]models.Refinement, len(records))
	for i := range records {
		refinements[i] = marshalRefinementDAO(records[i])
	}
	return refinements
}
