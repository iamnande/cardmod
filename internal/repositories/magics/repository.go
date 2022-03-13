package magics

import (
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
	"github.com/iamnande/cardmod/internal/repositories/errors"
)

// repository is the concrete implementation of the daos.MagicDAO interface.
type repository struct {
	magics []models.Magic
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.MagicDAO = (*repository)(nil)

// NewRepository initializes a new magic repository.
// TODO: sync/once this.
func NewRepository() repository {

	// new: initialize repository
	repo := repository{
		magics: make([]models.Magic, len(magics)),
	}

	// new: load the data into the model
	for i := 0; i < len(magics); i++ {
		repo.magics[i] = magics[i]
	}

	// new: return constructed repository instance
	return repo

}

// ListMagics gets a collection of magics.
func (r repository) ListMagics() []models.Magic {
	return r.magics
}

// GetMagic gets a magic by name.
func (r repository) GetMagic(name string) (models.Magic, error) {
	for i := 0; i < len(r.magics); i++ {
		if r.magics[i].Name() == name {
			return r.magics[i], nil
		}
	}
	return nil, errors.NewRepositoryError("magic not found")
}
