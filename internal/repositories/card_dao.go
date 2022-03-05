package repositories

import (
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/models"
)

// cardDAOWrapper is the concrete instance of a card.
type cardDAOWrapper struct {
	card *database.Card
}

func (db *cardDAOWrapper) Name() string {
	return db.card.Name
}

// marshalCardDAO creates a DAO instance from the database schema.
func marshalCardDAO(record *database.Card) models.Card {
	return &cardDAOWrapper{record}
}

// marshalCardDAO creates a DAO instance from the database schema.
func marshalCardDAOs(records []*database.Card) []models.Card {
	cards := make([]models.Card, len(records))
	for i := range records {
		cards[i] = marshalCardDAO(records[i])
	}
	return cards
}
