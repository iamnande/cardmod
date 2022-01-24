package calculationapi

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/pkg/api/calculationv1"
)

// api is the service implementation of the generate Calculation gRPC service.
type api struct {
	calculationv1.UnimplementedCalculationAPIServer
	calculationRepository daos.CalculationDAO
	cardRepository        daos.CardDAO
}

// New initializes a new card api instance.
func New(calculationRepository daos.CalculationDAO, cardRepository daos.CardDAO) api {
	return api{
		calculationRepository: calculationRepository,
		cardRepository:        cardRepository,
	}
}

// Calculate will calculate the number, and which type, of cards are needed for n number of a given magic.
func (api *api) Calculate(ctx context.Context,
	request *calculationv1.CalculateRequest) (*calculationv1.CalculateResponse, error) {

	// calculate: fetch (list) the calculations
	// NOTE: this is a list instead of a get to enable future options of
	// conversation (many cards satisfy the same magic).
	calculation, err := api.calculationRepository.ListCalculations(ctx, request.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch %s magic: %s", request.GetName(), err)
	}

	// calculate: fetch the card needed
	card, err := api.cardRepository.GetCard(ctx, calculation[0].CardID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch card details: %s", err)
	}

	// calculate: perform calculation
	// TODO: make this not so brittle with the slice[0] business
	count := request.GetCount() / calculation[0].MagicRatio() * calculation[0].CardRatio()

	// calculate: return response to caller
	return &calculationv1.CalculateResponse{
		Name:  card.Name(),
		Count: count,
	}, nil

}
