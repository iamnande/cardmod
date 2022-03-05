package itemv1

import (
	"context"

	"google.golang.org/grpc/codes"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
	"github.com/iamnande/cardmod/pkg/api/itemv1"
)

// api is the concrete implementation of the ItemAPI gRPC service.
type api struct {
	itemv1.UnimplementedItemAPIServer
	itemRepository daos.ItemDAO
}

// New initializes the v1 ItemAPI.
func New(repo daos.ItemDAO) *api {
	return &api{
		itemRepository: repo,
	}
}

// GetItem gets a item.
func (a *api) GetItem(ctx context.Context, req *itemv1.GetItemRequest) (*itemv1.Item, error) {

	// get: fetch the item
	res, err := a.itemRepository.GetItem(ctx, req.GetName())
	if err != nil {
		if _, ok := err.(*cerrors.APIError); ok {
			return nil, err
		}
		return nil, &cerrors.APIError{
			Code:      codes.Internal,
			Message:   "failed to get item",
			BaseError: err,
		}
	}

	// get: return the item
	return marshalItem(res), nil

}

// ListItems gets a collection of items.
func (a *api) ListItems(ctx context.Context, _ *itemv1.ListItemsRequest) (*itemv1.ListItemsResponse, error) {

	// list: fetch the items
	res, err := a.itemRepository.ListItems(ctx)
	if err != nil {
		return nil, &cerrors.APIError{
			Code:      codes.Internal,
			Message:   "failed to list items",
			BaseError: err,
		}
	}

	// list: return the item
	return &itemv1.ListItemsResponse{
		Items: marshalItems(res),
	}, nil

}

// marshalItem translates the DAO model to an API model.
func marshalItem(item models.Item) *itemv1.Item {
	return &itemv1.Item{
		Name: item.Name(),
	}
}

// marshalItems translates a collection of DAO models into a collection API model.
func marshalItems(items []models.Item) []*itemv1.Item {
	res := make([]*itemv1.Item, len(items))
	for i := 0; i < len(items); i++ {
		res[i] = marshalItem(items[i])
	}
	return res
}
