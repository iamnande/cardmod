package cardv1

import (
	"context"

	"google.golang.org/grpc/codes"

	"github.com/iamnande/cardmod/internal/api/errors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
	"github.com/iamnande/cardmod/internal/repositories/cards"
	"github.com/iamnande/cardmod/pkg/api/cardv1"
)

// api is the concrete implementation of the CardAPI gRPC service.
type api struct {
	cardv1.UnimplementedCardAPIServer
	cardRepository daos.CardDAO
}

// New initializes the v1 CardAPI.
func New() *api {
	return &api{
		cardRepository: cards.NewRepository(),
	}
}

// GetCard gets a card.
func (a *api) GetCard(_ context.Context, req *cardv1.GetCardRequest) (*cardv1.Card, error) {

	// get: fetch the card
	res, err := a.cardRepository.GetCard(req.GetName())
	if err != nil {
		return nil, errors.NewAPIError(codes.NotFound, "card not found", err)
	}

	// get: return the card
	return marshalCard(res), nil

}

// ListCards gets a collection of cards.
func (a *api) ListCards(_ context.Context, _ *cardv1.ListCardsRequest) (*cardv1.ListCardsResponse, error) {

	// list: return the cards
	return &cardv1.ListCardsResponse{
		Cards: marshalCards(a.cardRepository.ListCards()),
	}, nil

}

// marshalCard translates the DAO model to an API model.
func marshalCard(card models.Card) *cardv1.Card {
	return &cardv1.Card{
		Name:  card.Name(),
		Level: card.Level(),
	}
}

// marshalCards translates a collection of DAO models into a collection API model.
func marshalCards(cards []models.Card) []*cardv1.Card {
	res := make([]*cardv1.Card, len(cards))
	for i := 0; i < len(cards); i++ {
		res[i] = marshalCard(cards[i])
	}
	return res
}
