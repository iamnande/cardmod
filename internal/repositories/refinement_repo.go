package repositories

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/database/refinement"
	"github.com/iamnande/cardmod/internal/models"
)

// refinementRepository is the repository for managing refinements.
type refinementRepository struct {
	db *database.Client
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.RefinementDAO = (*refinementRepository)(nil)

// NewRefinementRepository initializes a new refinement repository instance.
func NewRefinementRepository(client *database.Client) *refinementRepository {
	return &refinementRepository{
		db: client,
	}
}

// CreateRefinement creates a new refinement.
func (r *refinementRepository) CreateRefinement(ctx context.Context,
	source, target string, numerator, denominator int32) (models.Refinement, error) {

	// create: initialize transaction
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	// create: create the refinement
	m, err := tx.Refinement.Create().
		SetSource(source).
		SetTarget(target).
		SetNumerator(numerator).
		SetDenominator(denominator).
		Save(ctx)
	if err != nil {
		if database.IsConstraintError(err) {
			return nil, rollback(tx, &cerrors.APIError{
				Code:      codes.AlreadyExists,
				Message:   "refinement already exists",
				BaseError: err,
			})
		}
		return nil, rollback(tx, err)
	}

	// create: return the newly created refinement
	return marshalRefinementDAO(m), tx.Commit()

}

// GetRefinement gets a refinement.
func (r *refinementRepository) GetRefinement(ctx context.Context, source, target string) (models.Refinement, error) {

	// get: gets the refinement
	refinement, err := r.db.Refinement.Query().Where(
		refinement.And(
			refinement.SourceEQ(source),
			refinement.TargetEQ(target),
		),
	).First(ctx)
	if err != nil {
		if database.IsNotFound(err) {
			return nil, &cerrors.APIError{
				Code:      codes.NotFound,
				Message:   "refinement not found",
				BaseError: err,
			}
		}
		return nil, err
	}

	// get: return the refinement to be retrieved
	return marshalRefinementDAO(refinement), nil

}

// ListRefinements gets a collection of refinements.
func (r *refinementRepository) ListRefinements(ctx context.Context) ([]models.Refinement, error) {

	// list: gets the refinements
	refinements, err := r.db.Refinement.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	// list: return the refinements to be retrieved
	return marshalRefinementDAOs(refinements), nil

}
