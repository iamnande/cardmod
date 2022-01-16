package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/domains/magic"
)

// magicRepository is a repository layer for accessing magic entities in the data layer.
type magicRepository struct {
	client *database.Client
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.MagicDAO = (*magicRepository)(nil)

// NewMagicRepository initializes a new magic repository instance.
func NewMagicRepository(client *database.Client) *magicRepository {
	return &magicRepository{
		client: client,
	}
}

// ListMagics lists all available magic entities.
func (repo *magicRepository) ListMagics(ctx context.Context) ([]magic.Magic, error) {

	// list: list the magics
	magics, err := repo.client.Magic.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	// list: return the list of magics
	return marshalMagicContainers(magics), nil

}

// CreateMagic creates a new magic entity.
func (repo *magicRepository) CreateMagic(ctx context.Context, name string) (magic.Magic, error) {

	// create: create the magic
	magic, err := repo.client.Magic.Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// create: return the newly created magic
	return &magicContainer{magic}, nil

}

// DescribeMagic describes a magic entity.
func (repo *magicRepository) DescribeMagic(ctx context.Context, id uuid.UUID) (magic.Magic, error) {

	// describe: describes the magic
	magic, err := repo.client.Magic.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// describe: return the magic to be described
	return &magicContainer{magic}, nil

}

// DeleteMagic deletes an existing magic entity.
func (repo *magicRepository) DeleteMagic(ctx context.Context, id uuid.UUID) error {

	// delete: delete the magic
	if err := repo.client.Magic.DeleteOneID(id).Exec(ctx); err != nil {
		return err
	}

	// delete: return "success"
	return nil

}
