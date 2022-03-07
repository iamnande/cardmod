package limitbreaks

import (
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
	"github.com/iamnande/cardmod/internal/repositories/errors"
)

// repository is the concrete implementation of the daos.LimitBreakDAO interface.
type repository struct {
	limitbreaks []models.LimitBreak
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.LimitBreakDAO = (*repository)(nil)

// NewRepository initializes a new limitbreak repository.
// TODO: sync/once this.
func NewRepository() repository {

	// new: initialize repository
	repo := repository{
		limitbreaks: make([]models.LimitBreak, len(limitbreaks)),
	}

	// new: load the data into the model
	for i := 0; i < len(limitbreaks); i++ {
		repo.limitbreaks[i] = limitbreaks[i]
	}

	// new: return constructed repository instance
	return repo

}

// ListLimitBreaks gets a collection of limitbreaks.
func (r repository) ListLimitBreaks() []models.LimitBreak {
	return r.limitbreaks
}

// GetLimitBreak gets a limitbreak by name.
func (r repository) GetLimitBreak(name string) (models.LimitBreak, error) {
	for i := 0; i < len(r.limitbreaks); i++ {
		if r.limitbreaks[i].Name() == name {
			return r.limitbreaks[i], nil
		}
	}
	return nil, errors.NewRepositoryError("limit break not found")
}
