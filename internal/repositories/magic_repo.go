package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/grpc/codes"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/database/magic"
	"github.com/iamnande/cardmod/internal/models"
)

// magicRepository is the repository for managing magics.
type magicRepository struct {
	db *database.Client
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.MagicDAO = (*magicRepository)(nil)

// NewMagicRepository initializes a new magic repository instance.
func NewMagicRepository(client *database.Client) *magicRepository {
	return &magicRepository{
		db: client,
	}
}

// CreateMagic creates a new magic.
func (r *magicRepository) CreateMagic(ctx context.Context, name string, purpose string) (models.Magic, error) {

	// create: determine purpose
	var magicPurpose magic.Purpose
	switch purpose {
	case magic.PurposeIndirect.String():
		magicPurpose = magic.PurposeIndirect
	case magic.PurposeOffensive.String():
		magicPurpose = magic.PurposeOffensive
	case magic.PurposeRestorative.String():
		magicPurpose = magic.PurposeRestorative
	default:
		return nil, &cerrors.APIError{
			Code:      codes.InvalidArgument,
			Message:   "invalid magic purpose",
			BaseError: fmt.Errorf("purpose %s is invalid", purpose),
		}
	}

	// create: initialize transaction
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	// create: create the magic
	m, err := tx.Magic.Create().
		SetName(name).
		SetPurpose(magicPurpose).
		Save(ctx)
	if err != nil {
		if database.IsConstraintError(err) {
			return nil, rollback(tx, &cerrors.APIError{
				Code:      codes.AlreadyExists,
				Message:   "magic already exists",
				BaseError: err,
			})
		}
		return nil, rollback(tx, err)
	}

	// create: return the newly created magic
	return marshalMagicDAO(m), tx.Commit()

}

// GetMagic gets a magic.
func (r *magicRepository) GetMagic(ctx context.Context, name string) (models.Magic, error) {

	// get: gets the magic
	magic, err := r.db.Magic.Query().Where(magic.NameEQ(name)).First(ctx)
	if err != nil {
		if database.IsNotFound(err) {
			return nil, &cerrors.APIError{
				Code:      codes.NotFound,
				Message:   "magic not found",
				BaseError: err,
			}
		}
		return nil, err
	}

	// get: return the magic to be retrieved
	return marshalMagicDAO(magic), nil

}

// ListMagics gets a collection of magics.
func (r *magicRepository) ListMagics(ctx context.Context) ([]models.Magic, error) {

	// list: gets the magics
	magics, err := r.db.Magic.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	// list: return the magics to be retrieved
	return marshalMagicDAOs(magics), nil

}
