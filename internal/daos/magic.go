package daos

import (
	"context"

	"github.com/google/uuid"

	"github.com/iamnande/cardmod/internal/domains/magic"
)

// MagicDAO is the data access object (controller<->data layer) interface definition.
type MagicDAO interface {

	// ListMagics lists all magic entities.
	ListMagics(ctx context.Context) ([]magic.Magic, error)

	// CreateMagic creates a magic entity.
	CreateMagic(ctx context.Context, name string) (magic.Magic, error)

	// DescribeMagic describes a magic entity.
	DescribeMagic(ctx context.Context, id uuid.UUID) (magic.Magic, error)

	// DeleteMagic deletes a magic entity.
	DeleteMagic(ctx context.Context, id uuid.UUID) error
}
