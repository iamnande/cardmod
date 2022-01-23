package card

import (
	"github.com/google/uuid"
)

// Card is the domain interface to be used as a data access object translator between the data layer and the Card repository.
//go:generate mockgen -source card.go -destination=./mocks/card.go -package mocks
type Card interface {
	ID() uuid.UUID
	Name() string
}
