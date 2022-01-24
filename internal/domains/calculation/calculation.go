package calculation

import (
	"github.com/google/uuid"
)

// Calculation is the domain interface to be used as a data access
// object translator between the data layer and the Calculation repository.
//go:generate mockgen -source calculation.go -destination=./mocks/calculation.go -package mocks
type Calculation interface {
	ID() uuid.UUID
	CardID() uuid.UUID
	MagicID() uuid.UUID
	CardRatio() int32
	MagicRatio() int32
}
