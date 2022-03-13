package refinementv1

import (
	"context"

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

// ListRefinements gets a collection of refinements.
func (a *api) ListRefinements(_ context.Context,
	_ *refinementv1.ListRefinementsRequest) (*refinementv1.ListRefinementsResponse, error) {

	// list: fetch filtered list of refinements
	res := a.refinementRepository.ListRefinements(nil)

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
