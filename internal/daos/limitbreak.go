package daos

import (
	"context"

	"github.com/iamnande/cardmod/internal/models"
)

// LimitBreakDAO is the DAO for the LimitBreak model.
//go:generate mockgen -source limitbreak.go -destination=./mocks/limitbreak.go -package mocks
type LimitBreakDAO interface {

	// Gets a limit break.
	GetLimitBreak(ctx context.Context, name string) (models.LimitBreak, error)

	// Lists a collection of limit breaks.
	ListLimitBreaks(ctx context.Context) ([]models.LimitBreak, error)
}
