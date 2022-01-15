package magicapi

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamnande/cardmod/pkg/api/magicv1"
)

// api is the service implementation of the generate MagicAPI gRPC service.
type api struct {
	magicv1.UnimplementedMagicAPIServer
}

// New initializes a new magic api instance.
func New() api {
	return api{}
}

// CreateMagic creates a new magic entity.
func (api *api) CreateMagic(ctx context.Context, request *magicv1.CreateMagicRequest) (*magicv1.CreateMagicResponse, error) {

	// create: return response to caller
	return &magicv1.CreateMagicResponse{
		Magic: &magicv1.Magic{
			Id:   uuid.New().String(),
			Name: request.GetName(),
		},
	}, nil

}

// DescribeMagic describes an existing magic entity.
func (api *api) DescribeMagic(ctx context.Context, request *magicv1.DescribeMagicRequest) (*magicv1.DescribeMagicResponse, error) {

	// create: return response to caller
	return &magicv1.DescribeMagicResponse{
		Magic: &magicv1.Magic{
			Id:   uuid.New().String(),
			Name: "Fastitocalon",
		},
	}, nil

}
