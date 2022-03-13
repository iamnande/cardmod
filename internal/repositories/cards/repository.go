package cards

import (
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
	"github.com/iamnande/cardmod/internal/repositories/errors"
)

// repository is the concrete implementation of the daos.CardDAO interface.
type repository struct {
	cards []models.Card
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.CardDAO = (*repository)(nil)

// NewRepository initializes a new card repository.
// TODO: sync/once this.
func NewRepository() repository {

	// new: initialize repository
	repo := repository{
		cards: make([]models.Card, len(cards)),
	}

	// new: load the data into the model
	for i := 0; i < len(cards); i++ {
		repo.cards[i] = cards[i]
	}

	// new: return constructed repository instance
	return repo

}

// ListCards gets a collection of cards.
func (r repository) ListCards() []models.Card {
	return r.cards
}

// GetCard gets a card by name.
func (r repository) GetCard(name string) (models.Card, error) {
	for i := 0; i < len(r.cards); i++ {
		if r.cards[i].Name() == name {
			return r.cards[i], nil
		}
	}
	return nil, errors.NewRepositoryError("card not found")
}
