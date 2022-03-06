package refinementv1

import (
	"context"

	"google.golang.org/grpc/codes"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
	"github.com/iamnande/cardmod/pkg/api/refinementv1"
)

// api is the concrete implementation of the RefinementAPI gRPC service.
type api struct {
	refinementv1.UnimplementedRefinementAPIServer
	refinementRepository daos.RefinementDAO
}

// New initializes the v1 RefinementAPI.
func New(repo daos.RefinementDAO) *api {
	return &api{
		refinementRepository: repo,
	}
}

// GetRefinement gets a refinement.
func (a *api) GetRefinement(ctx context.Context, req *refinementv1.GetRefinementRequest) (*refinementv1.Refinement, error) {

	// get: fetch the refinement
	res, err := a.refinementRepository.GetRefinement(ctx, req.GetSource(), req.GetTarget())
	if err != nil {
		if _, ok := err.(*cerrors.APIError); ok {
			return nil, err
		}
		return nil, &cerrors.APIError{
			Code:      codes.Internal,
			Message:   "failed to get refinement",
			BaseError: err,
		}
	}

	// get: return the refinement
	return marshalRefinement(res), nil

}

// ListRefinements gets a collection of refinements.
func (a *api) ListRefinements(ctx context.Context,
	_ *refinementv1.ListRefinementsRequest) (*refinementv1.ListRefinementsResponse, error) {

	// list: fetch the refinements
	res, err := a.refinementRepository.ListRefinements(ctx)
	if err != nil {
		return nil, &cerrors.APIError{
			Code:      codes.Internal,
			Message:   "failed to list refinements",
			BaseError: err,
		}
	}

	// list: return the refinement
	return &refinementv1.ListRefinementsResponse{
		Refinements: marshalRefinements(res),
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
