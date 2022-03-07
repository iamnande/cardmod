package daos

import (
	"github.com/iamnande/cardmod/internal/models"
)

// CardDAO is the DAO for the Card model.
//go:generate mockgen -source card.go -destination=./mocks/card.go -package mocks
type CardDAO interface {

	// Lists a collection of cards.
	ListCards() []models.Card

	// Gets a card.
	GetCard(name string) (models.Card, error)
}
