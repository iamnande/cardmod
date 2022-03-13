package daos

import (
	"github.com/iamnande/cardmod/internal/models"
)

// RefinementDAO is the DAO for the Refinement model.
type RefinementDAO interface {

	// Lists a collection of refinements.
	ListRefinements(filter models.RefinementFilter) []models.Refinement
}
