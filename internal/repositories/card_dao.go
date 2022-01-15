package repositories

import (
	"github.com/google/uuid"

	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/domains/card"
)

// cardContainer is a wrapper around the data layer schema that satisfies the domain model.
type cardContainer struct {
	card *database.Card
}

func (db *cardContainer) ID() uuid.UUID {
	return db.card.ID
}

func (db *cardContainer) Name() string {
	return db.card.Name
}

// marshalCardContainers creates containers from the data layer schema.
func marshalCardContainers(records []*database.Card) []card.Card {
	cards := make([]card.Card, len(records))
	for i := range records {
		cards[i] = &cardContainer{records[i]}
	}
	return cards
}
