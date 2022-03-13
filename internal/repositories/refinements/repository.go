package refinements

import (
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
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
// Optionally you can filter by source, target, or both.
func (r repository) ListRefinements(filter models.RefinementFilter) []models.Refinement {

	// list: if no filter, return complete collection
	if filter == nil {
		return r.refinements
	}

	// list: construct empty return slice
	res := make([]models.Refinement, 0)

	// list: iterate the refinements, filtering as necessary
	for i := 0; i < len(r.refinements); i++ {

		// list: filter on just target
		if filter.Target() != "" && filter.Source() == "" {
			if r.refinements[i].Target() == filter.Target() {
				res = append(res, refinements[i])
				continue
			}
		}

		// list: filter on just source
		if filter.Source() != "" && filter.Target() == "" {
			if r.refinements[i].Source() == filter.Source() {
				res = append(res, refinements[i])
				continue
			}
		}

		// list: filter on both source and target
		if filter.Source() != "" && filter.Target() != "" {
			if r.refinements[i].Source() == filter.Source() &&
				r.refinements[i].Target() == filter.Target() {
				res = append(res, refinements[i])
				continue
			}
		}

	}

	// list: return filtered refinements
	return res

}
