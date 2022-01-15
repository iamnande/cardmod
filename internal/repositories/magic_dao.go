package repositories

import (
	"github.com/google/uuid"

	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/domains/magic"
)

// magicContainer is a wrapper around the data layer schema that satisfies the domain model.
type magicContainer struct {
	magic *database.Magic
}

func (db *magicContainer) ID() uuid.UUID {
	return db.magic.ID
}

func (db *magicContainer) Name() string {
	return db.magic.Name
}

// marshalMagicContainers creates containers from the data layer schema.
func marshalMagicContainers(records []*database.Magic) []magic.Magic {
	cards := make([]magic.Magic, len(records))
	for i := range records {
		cards[i] = &magicContainer{records[i]}
	}
	return cards
}
