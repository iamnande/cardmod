package cardapi

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/domains/card"
	"github.com/iamnande/cardmod/pkg/api/cardv1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// api is the service implementation of the generate CardAPI gRPC service.
type api struct {
	cardv1.UnimplementedCardAPIServer
	cardRepository daos.CardDAO
}

// New initializes a new card api instance.
func New(cardRepository daos.CardDAO) api {
	return api{
		cardRepository: cardRepository,
	}
}

// ListCards lists all available card entities.
func (api *api) ListCards(ctx context.Context, request *cardv1.ListCardsRequest) (*cardv1.ListCardsResponse, error) {

	// list: list the available cards
	cards, err := api.cardRepository.ListCards(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list cards: %s", err)
	}

	// list: return response to caller
	return &cardv1.ListCardsResponse{
		Cards: marshalCards(cards),
	}, nil

}

// DescribeCard describes an existing card entity.
func (api *api) DescribeCard(ctx context.Context, request *cardv1.DescribeCardRequest) (*cardv1.DescribeCardResponse, error) {

	// describe: parse input uuid
	id, err := uuid.Parse(request.GetCardId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid card id")
	}

	// describe: describe card
	card, err := api.cardRepository.DescribeCard(ctx, id)
	if err != nil {
		return nil, err
	}

	// describe: return card to caller
	return &cardv1.DescribeCardResponse{
		Card: marshalCard(card),
	}, nil

}

func marshalCard(card card.Card) *cardv1.Card {
	return &cardv1.Card{
		Name: card.Name(),
		Id:   card.ID().String(),
	}
}

func marshalCards(records []card.Card) []*cardv1.Card {
	cards := make([]*cardv1.Card, len(records))
	for i := range cards {
		cards[i] = marshalCard(records[i])
	}
	return cards
}
