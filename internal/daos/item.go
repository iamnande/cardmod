package daos

import (
	"context"

	"github.com/iamnande/cardmod/internal/models"
)

// ItemDAO is the DAO for the Item model.
//go:generate mockgen -source item.go -destination=./mocks/item.go -package mocks
type ItemDAO interface {

	// Gets a item.
	GetItem(ctx context.Context, name string) (models.Item, error)

	// Lists a collection of items.
	ListItems(ctx context.Context) ([]models.Item, error)
}
