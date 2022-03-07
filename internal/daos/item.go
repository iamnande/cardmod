package daos

import (
	"github.com/iamnande/cardmod/internal/models"
)

// ItemDAO is the DAO for the Item model.
//go:generate mockgen -source item.go -destination=./mocks/item.go -package mocks
type ItemDAO interface {

	// Lists a collection of items.
	ListItems() []models.Item

	// Gets a item.
	GetItem(name string) (models.Item, error)
}
