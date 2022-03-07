package daos

import (
	"github.com/iamnande/cardmod/internal/models"
)

// MagicDAO is the DAO for the Magic model.
//go:generate mockgen -source magic.go -destination=./mocks/magic.go -package mocks
type MagicDAO interface {

	// Lists a collection of magics.
	ListMagics() []models.Magic

	// Gets a magic.
	GetMagic(name string) (models.Magic, error)
}
