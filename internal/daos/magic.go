package daos

import (
	"github.com/iamnande/cardmod/internal/models"
)

// MagicDAO is the DAO for the Magic model.
type MagicDAO interface {

	// Lists a collection of magics.
	ListMagics() []models.Magic

	// Gets a magic.
	GetMagic(name string) (models.Magic, error)
}
