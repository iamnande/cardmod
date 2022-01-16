package card

import (
	"github.com/google/uuid"
)

// Card is the domain interface to be used as a data access object translator between the data layer and the Card repository.
type Card interface {
	ID() uuid.UUID
	Name() string
}
