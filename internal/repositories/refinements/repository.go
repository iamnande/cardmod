package refinements

import (
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
	"github.com/iamnande/cardmod/internal/repositories/errors"
)

// repository is the concrete implementation of the daos.RefinementDAO interface.
type repository struct {
	refinements []models.Refinement
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.RefinementDAO = (*repository)(nil)

// NewRepository initializes a new refinement repository.
// TODO: sync/once this.
func NewRepository() repository {

	// new: initialize repository
	repo := repository{
		refinements: make([]models.Refinement, len(refinements)),
	}

	// new: load the data into the model
	for i := 0; i < len(refinements); i++ {
		repo.refinements[i] = refinements[i]
	}

	// new: return constructed repository instance
	return repo

}

// ListRefinements gets a collection of refinements.
func (r repository) ListRefinements() []models.Refinement {
	return r.refinements
}

// GetRefinement gets a refinement by source and target.
func (r repository) GetRefinement(source, target string) (models.Refinement, error) {
	for i := 0; i < len(r.refinements); i++ {
		if r.refinements[i].Source() == source &&
			r.refinements[i].Target() == target {
			return r.refinements[i], nil
		}
	}
	return nil, errors.NewRepositoryError("refinement not found")
}
