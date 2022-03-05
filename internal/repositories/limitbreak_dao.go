package repositories

import (
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/models"
)

// limitbreakDAOWrapper is the concrete instance of a limitbreak.
type limitbreakDAOWrapper struct {
	limitbreak *database.LimitBreak
}

func (db *limitbreakDAOWrapper) Name() string {
	return db.limitbreak.Name
}

// marshalLimitBreakDAO creates a DAO instance from the database schema.
func marshalLimitBreakDAO(record *database.LimitBreak) models.LimitBreak {
	return &limitbreakDAOWrapper{record}
}

// marshalLimitBreakDAO creates a DAO instance from the database schema.
func marshalLimitBreakDAOs(records []*database.LimitBreak) []models.LimitBreak {
	limitbreaks := make([]models.LimitBreak, len(records))
	for i := range records {
		limitbreaks[i] = marshalLimitBreakDAO(records[i])
	}
	return limitbreaks
}
