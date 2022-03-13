package magicv1

import (
	"context"

	"google.golang.org/grpc/codes"

	"github.com/iamnande/cardmod/internal/api/errors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
	"github.com/iamnande/cardmod/internal/repositories/magics"
	"github.com/iamnande/cardmod/pkg/api/magicv1"
)

// api is the concrete implementation of the MagicAPI gRPC service.
type api struct {
	magicv1.UnimplementedMagicAPIServer
	magicRepository daos.MagicDAO
}

// New initializes the v1 MagicAPI.
func New() *api {
	return &api{
		magicRepository: magics.NewRepository(),
	}
}

// GetMagic gets a magic.
func (a *api) GetMagic(_ context.Context, req *magicv1.GetMagicRequest) (*magicv1.Magic, error) {

	// get: fetch the magic
	res, err := a.magicRepository.GetMagic(req.GetName())
	if err != nil {
		return nil, errors.NewAPIError(codes.NotFound, "magic not found", err)
	}

	// get: return the magic
	return marshalMagic(res), nil

}

// ListMagics gets a collection of magics.
func (a *api) ListMagics(_ context.Context, _ *magicv1.ListMagicsRequest) (*magicv1.ListMagicsResponse, error) {

	// list: return the magics
	return &magicv1.ListMagicsResponse{
		Magics: marshalMagics(a.magicRepository.ListMagics()),
	}, nil

}

// marshalMagic translates the DAO model to an API model.
func marshalMagic(magic models.Magic) *magicv1.Magic {
	return &magicv1.Magic{
		Name:    magic.Name(),
		Purpose: magic.Purpose(),
	}
}

// marshalMagics translates a collection of DAO models into a collection API model.
func marshalMagics(magics []models.Magic) []*magicv1.Magic {
	res := make([]*magicv1.Magic, len(magics))
	for i := 0; i < len(magics); i++ {
		res[i] = marshalMagic(magics[i])
	}
	return res
}
