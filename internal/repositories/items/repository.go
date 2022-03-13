package items

import (
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
	"github.com/iamnande/cardmod/internal/repositories/errors"
)

// repository is the concrete implementation of the daos.ItemDAO interface.
type repository struct {
	items []models.Item
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.ItemDAO = (*repository)(nil)

// NewRepository initializes a new item repository.
// TODO: sync/once this.
func NewRepository() repository {

	// new: initialize repository
	repo := repository{
		items: make([]models.Item, len(items)),
	}

	// new: load the data into the model
	for i := 0; i < len(items); i++ {
		repo.items[i] = items[i]
	}

	// new: return constructed repository instance
	return repo

}

// ListItems gets a collection of items.
func (r repository) ListItems() []models.Item {
	return r.items
}

// GetItem gets a item by name.
func (r repository) GetItem(name string) (models.Item, error) {
	for i := 0; i < len(r.items); i++ {
		if r.items[i].Name() == name {
			return r.items[i], nil
		}
	}
	return nil, errors.NewRepositoryError("item not found")
}
