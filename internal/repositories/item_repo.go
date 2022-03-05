package repositories

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/database/item"
	"github.com/iamnande/cardmod/internal/models"
)

// itemRepository is the repository for managing items.
type itemRepository struct {
	db *database.Client
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.ItemDAO = (*itemRepository)(nil)

// NewItemRepository initializes a new item repository instance.
func NewItemRepository(client *database.Client) *itemRepository {
	return &itemRepository{
		db: client,
	}
}

// CreateItem creates a new item.
func (r *itemRepository) CreateItem(ctx context.Context, name string) (models.Item, error) {

	// create: initialize transaction
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	// create: create the item
	item, err := tx.Item.Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		if database.IsConstraintError(err) {
			return nil, rollback(tx, &cerrors.APIError{
				Code:      codes.AlreadyExists,
				Message:   "item already exists",
				BaseError: err,
			})
		}
		return nil, rollback(tx, err)
	}

	// create: return the newly created item
	return marshalItemDAO(item), tx.Commit()

}

// GetItem gets a item.
func (r *itemRepository) GetItem(ctx context.Context, name string) (models.Item, error) {

	// get: gets the item
	item, err := r.db.Item.Query().Where(item.NameEQ(name)).First(ctx)
	if err != nil {
		if database.IsNotFound(err) {
			return nil, &cerrors.APIError{
				Code:      codes.NotFound,
				Message:   "item not found",
				BaseError: err,
			}
		}
		return nil, err
	}

	// get: return the item to be retrieved
	return marshalItemDAO(item), nil

}

// ListItems gets a collection of items.
func (r *itemRepository) ListItems(ctx context.Context) ([]models.Item, error) {

	// list: gets the items
	items, err := r.db.Item.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	// list: return the items to be retrieved
	return marshalItemDAOs(items), nil

}
