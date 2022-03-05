package repositories

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/database/card"
	"github.com/iamnande/cardmod/internal/models"
)

// cardRepository is the repository for managing cards.
type cardRepository struct {
	db *database.Client
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.CardDAO = (*cardRepository)(nil)

// NewCardRepository initializes a new card repository instance.
// TODO: repository specific errors (think APIError, but RepositoryError).
func NewCardRepository(client *database.Client) *cardRepository {
	return &cardRepository{
		db: client,
	}
}

// CreateCard creates a new card.
func (r *cardRepository) CreateCard(ctx context.Context, name string, level int32) (models.Card, error) {

	// create: initialize transaction
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	// create: create the card
	card, err := tx.Card.Create().
		SetName(name).
		SetLevel(level).
		Save(ctx)
	if err != nil {
		if database.IsConstraintError(err) {
			return nil, rollback(tx, &cerrors.APIError{
				Code:      codes.AlreadyExists,
				Message:   "card already exists",
				BaseError: err,
			})
		}
		return nil, rollback(tx, err)
	}

	// create: return the newly created card
	return marshalCardDAO(card), tx.Commit()

}

// GetCard gets a card.
func (r *cardRepository) GetCard(ctx context.Context, name string) (models.Card, error) {

	// get: gets the card
	card, err := r.db.Card.Query().Where(card.NameEQ(name)).First(ctx)
	if err != nil {
		if database.IsNotFound(err) {
			return nil, &cerrors.APIError{
				Code:      codes.NotFound,
				Message:   "card not found",
				BaseError: err,
			}
		}
		return nil, err
	}

	// get: return the card to be retrieved
	return marshalCardDAO(card), nil

}

// ListCards gets a collection of cards.
func (r *cardRepository) ListCards(ctx context.Context) ([]models.Card, error) {

	// list: gets the cards
	cards, err := r.db.Card.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	// list: return the cards to be retrieved
	return marshalCardDAOs(cards), nil

}
