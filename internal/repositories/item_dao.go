package repositories

import (
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/models"
)

// itemDAOWrapper is the concrete instance of a item.
type itemDAOWrapper struct {
	item *database.Item
}

func (db *itemDAOWrapper) Name() string {
	return db.item.Name
}

// marshalItemDAO creates a DAO instance from the database schema.
func marshalItemDAO(record *database.Item) models.Item {
	return &itemDAOWrapper{record}
}

// marshalItemDAO creates a DAO instance from the database schema.
func marshalItemDAOs(records []*database.Item) []models.Item {
	items := make([]models.Item, len(records))
	for i := range records {
		items[i] = marshalItemDAO(records[i])
	}
	return items
}
