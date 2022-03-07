package daos

import (
	"github.com/iamnande/cardmod/internal/models"
)

// LimitBreakDAO is the DAO for the LimitBreak model.
type LimitBreakDAO interface {

	// Lists a collection of limit breaks.
	ListLimitBreaks() []models.LimitBreak

	// Gets a limit break.
	GetLimitBreak(name string) (models.LimitBreak, error)
}
