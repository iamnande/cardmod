package refinementv1

import (
	"context"

	"google.golang.org/grpc/codes"

	"github.com/iamnande/cardmod/internal/api/errors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
	"github.com/iamnande/cardmod/internal/repositories/refinements"
	"github.com/iamnande/cardmod/pkg/api/refinementv1"
)

// api is the concrete implementation of the RefinementAPI gRPC service.
type api struct {
	refinementv1.UnimplementedRefinementAPIServer
	refinementRepository daos.RefinementDAO
}

// New initializes the v1 RefinementAPI.
func New() *api {
	return &api{
		refinementRepository: refinements.NewRepository(),
	}
}

// GetRefinement gets a refinement.
func (a *api) GetRefinement(_ context.Context, req *refinementv1.GetRefinementRequest) (*refinementv1.Refinement, error) {

	// get: fetch the refinement
	res, err := a.refinementRepository.GetRefinement(req.GetSource(), req.GetTarget())
	if err != nil {
		return nil, errors.NewAPIError(codes.NotFound, "refinement not found", err)
	}

	// get: return the refinement
	return marshalRefinement(res), nil

}

// ListRefinements gets a collection of refinements.
func (a *api) ListRefinements(_ context.Context,
	_ *refinementv1.ListRefinementsRequest) (*refinementv1.ListRefinementsResponse, error) {

	// list: return the refinement
	return &refinementv1.ListRefinementsResponse{
		Refinements: marshalRefinements(a.refinementRepository.ListRefinements()),
	}, nil

}

// marshalRefinement translates the DAO model to an API model.
func marshalRefinement(refinement models.Refinement) *refinementv1.Refinement {
	return &refinementv1.Refinement{
		Source:      refinement.Source(),
		Target:      refinement.Target(),
		Numerator:   refinement.Numerator(),
		Denominator: refinement.Denominator(),
	}
}

// marshalRefinements translates a collection of DAO models into a collection API model.
func marshalRefinements(refinements []models.Refinement) []*refinementv1.Refinement {
	res := make([]*refinementv1.Refinement, len(refinements))
	for i := 0; i < len(refinements); i++ {
		res[i] = marshalRefinement(refinements[i])
	}
	return res
}
