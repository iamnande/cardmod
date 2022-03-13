package calculatev1

import (
	"context"
	"fmt"

	"github.com/iamnande/cardmod/internal/api/errors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/repositories/refinements"
	"github.com/iamnande/cardmod/pkg/api/calculatev1"
	"google.golang.org/grpc/codes"
)

// api is the concrete implementation of the CalculateAPI gRPC service.
type api struct {
	calculatev1.UnimplementedCalculateAPIServer
	refinementRepository daos.RefinementDAO
}

// this is a build/compile time check to ensure our concrete implementation satisfies the API interface.
var _ calculatev1.CalculateAPIServer = (*api)(nil)

// New initializes the v1 CalculateAPI.
func New() *api {
	return &api{
		refinementRepository: refinements.NewRepository(),
	}
}

// Calculates quantities and sources of a given refinement.
func (api *api) Calculate(_ context.Context, req *calculatev1.CalculateRequest) (*calculatev1.CalculateResponse, error) {

	// calculate: constructe empty response container
	res := make([]*calculatev1.CalculateResponse_Refinement, 0)

	// calculate: fetch the list of refinements
	collection := api.refinementRepository.ListRefinements(refinements.NewFilter("", req.GetTarget()))

	// calculate: not found check
	if len(collection) == 0 {
		return nil, errors.NewAPIError(codes.NotFound, "target not found",
			fmt.Errorf("%s target does not exist", req.GetTarget()))
	}

	// calculate: add found refinements to response collection
	for i := 0; i < len(collection); i++ {

		// calculate: perform the calculation based on request count
		count := req.GetCount() / collection[i].Denominator() * collection[i].Numerator()

		// calculate: add the refinement to the response array
		refinement := &calculatev1.CalculateResponse_Refinement{
			Source:      collection[i].Source(),
			Count:       count,
			Refinements: make([]*calculatev1.CalculateResponse_Refinement, 0),
		}

		// calculate: do the same thing for child refinements (if applicable)
		childCollection := api.refinementRepository.ListRefinements(refinements.NewFilter("", collection[i].Source()))
		for j := 0; j < len(childCollection); j++ {
			childCount := count * childCollection[j].Numerator()
			refinement.Refinements = append(refinement.Refinements, &calculatev1.CalculateResponse_Refinement{
				Source: childCollection[j].Source(),
				Count:  childCount,
			})
		}

		// calculate: add calculation tree to response
		res = append(res, refinement)

	}

	// calculate: return response
	return &calculatev1.CalculateResponse{
		Refinements: res,
	}, nil

}
