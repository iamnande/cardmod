package cardapi

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamnande/cardmod/pkg/api/cardv1"
)

// api is the service implementation of the generate CardAPI gRPC service.
type api struct {
	cardv1.UnimplementedCardAPIServer
}

// New initializes a new card api instance.
func New() api {
	return api{}
}

// CreateCard creates a new card entity.
func (api *api) CreateCard(ctx context.Context, request *cardv1.CreateCardRequest) (*cardv1.CreateCardResponse, error) {

	// create: return response to caller
	return &cardv1.CreateCardResponse{
		Card: &cardv1.Card{
			Id:   uuid.New().String(),
			Name: request.GetName(),
		},
	}, nil

}

// DescribeCard describes an existing card entity.
func (api *api) DescribeCard(ctx context.Context, request *cardv1.DescribeCardRequest) (*cardv1.DescribeCardResponse, error) {

	// create: return response to caller
	return &cardv1.DescribeCardResponse{
		Card: &cardv1.Card{
			Id:   uuid.New().String(),
			Name: "Fastitocalon",
		},
	}, nil

}
