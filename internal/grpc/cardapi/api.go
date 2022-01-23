package cardapi

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/domains/card"
	"github.com/iamnande/cardmod/pkg/api/cardv1"
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

	// list: if a search query is passed, use it
	query := ""
	if request.Filter != nil {
		if request.Filter.Query != "" {
			query = request.Filter.Query
		}
	}

	// list: list the available cards
	cards, err := api.cardRepository.ListCards(ctx, query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list cards: %s", err)
	}

	// list: return response to caller
	return &cardv1.ListCardsResponse{
		Cards: marshalCards(cards),
	}, nil

}

// GetCard gets an existing card entity.
func (api *api) GetCard(ctx context.Context, request *cardv1.GetCardRequest) (*cardv1.Card, error) {

	// get: parse input uuid
	id, err := uuid.Parse(request.GetCardId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid card id")
	}

	// get: get card
	card, err := api.cardRepository.GetCard(ctx, id)
	if err != nil {
		return nil, err
	}

	// get: return card to caller
	return marshalCard(card), nil

}

func marshalCard(card card.Card) *cardv1.Card {
	return &cardv1.Card{
		Name: card.Name(),
		Id:   card.ID().String(),
	}
}

func marshalCards(records []card.Card) []*cardv1.Card {
	cards := make([]*cardv1.Card, len(records))
	for i := 0; i < len(records); i++ {
		cards[i] = marshalCard(records[i])
	}
	return cards
}
