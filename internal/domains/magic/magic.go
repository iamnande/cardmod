package magic

import (
	"github.com/google/uuid"
)

// Magic is the domain interface to be used as a data access object translator between the data layer and the Magic repository.
type Magic interface {
	ID() uuid.UUID
	Name() string
}
