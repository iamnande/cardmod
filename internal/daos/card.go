package daos

import (
	"context"

	"github.com/iamnande/cardmod/internal/models"
)

// CardDAO is the DAO for the Card model.
//go:generate mockgen -source card.go -destination=./mocks/card.go -package mocks
type CardDAO interface {

	// Gets a card.
	GetCard(ctx context.Context, name string) (models.Card, error)

	// Lists a collection of cards.
	ListCards(ctx context.Context) ([]models.Card, error)
}
