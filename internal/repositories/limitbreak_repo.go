package repositories

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/database/limitbreak"
	"github.com/iamnande/cardmod/internal/models"
)

// limitbreakRepository is the repository for managing limitbreaks.
type limitbreakRepository struct {
	db *database.Client
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.LimitBreakDAO = (*limitbreakRepository)(nil)

// NewLimitBreakRepository initializes a new limitbreak repository instance.
// TODO: repository specific errors (think APIError, but RepositoryError).
func NewLimitBreakRepository(client *database.Client) *limitbreakRepository {
	return &limitbreakRepository{
		db: client,
	}
}

// CreateLimitBreak creates a new limitbreak.
func (r *limitbreakRepository) CreateLimitBreak(ctx context.Context, name string) (models.LimitBreak, error) {

	// create: initialize transaction
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	// create: create the limitbreak
	limitbreak, err := tx.LimitBreak.Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		if database.IsConstraintError(err) {
			return nil, rollback(tx, &cerrors.APIError{
				Code:      codes.AlreadyExists,
				Message:   "limit break already exists",
				BaseError: err,
			})
		}
		return nil, rollback(tx, err)
	}

	// create: return the newly created limitbreak
	return marshalLimitBreakDAO(limitbreak), tx.Commit()

}

// GetLimitBreak gets a limitbreak.
func (r *limitbreakRepository) GetLimitBreak(ctx context.Context, name string) (models.LimitBreak, error) {

	// get: gets the limitbreak
	limitbreak, err := r.db.LimitBreak.Query().Where(limitbreak.NameEQ(name)).First(ctx)
	if err != nil {
		if database.IsNotFound(err) {
			return nil, &cerrors.APIError{
				Code:      codes.NotFound,
				Message:   "limit break not found",
				BaseError: err,
			}
		}
		return nil, err
	}

	// get: return the limitbreak to be retrieved
	return marshalLimitBreakDAO(limitbreak), nil

}

// ListLimitBreaks gets a collection of limitbreaks.
func (r *limitbreakRepository) ListLimitBreaks(ctx context.Context) ([]models.LimitBreak, error) {

	// list: gets the limitbreaks
	limitbreaks, err := r.db.LimitBreak.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	// list: return the limitbreaks to be retrieved
	return marshalLimitBreakDAOs(limitbreaks), nil

}
