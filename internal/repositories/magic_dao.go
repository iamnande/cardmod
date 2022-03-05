package repositories

import (
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/models"
)

// magicDAOWrapper is the concrete instance of a magic.
type magicDAOWrapper struct {
	magic *database.Magic
}

func (db *magicDAOWrapper) Name() string {
	return db.magic.Name
}

// marshalMagicDAO creates a DAO instance from the database schema.
func marshalMagicDAO(record *database.Magic) models.Magic {
	return &magicDAOWrapper{record}
}

// marshalMagicDAO creates a DAO instance from the database schema.
func marshalMagicDAOs(records []*database.Magic) []models.Magic {
	magics := make([]models.Magic, len(records))
	for i := range records {
		magics[i] = marshalMagicDAO(records[i])
	}
	return magics
}
