package daos

import (
	"context"

	"github.com/iamnande/cardmod/internal/models"
)

// RefinementDAO is the DAO for the Refinement model.
//go:generate mockgen -source refinement.go -destination=./mocks/refinement.go -package mocks
type RefinementDAO interface {

	// Gets a refinement.
	GetRefinement(ctx context.Context, source, target string) (models.Refinement, error)

	// Lists a collection of refinements.
	ListRefinements(ctx context.Context) ([]models.Refinement, error)
}
