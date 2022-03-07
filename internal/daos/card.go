package daos

import (
	"github.com/iamnande/cardmod/internal/models"
)

// CardDAO is the DAO for the Card model.
type CardDAO interface {

	// Lists a collection of cards.
	ListCards() []models.Card

	// Gets a card.
	GetCard(name string) (models.Card, error)
}
