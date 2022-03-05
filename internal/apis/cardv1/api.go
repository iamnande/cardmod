package cardv1

import (
	"context"

	"google.golang.org/grpc/codes"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
	"github.com/iamnande/cardmod/pkg/api/cardv1"
)

// api is the concrete implementation of the CardAPI gRPC service.
type api struct {
	cardv1.UnimplementedCardAPIServer
	cardRepository daos.CardDAO
}

// New initializes the v1 CardAPI.
func New(repo daos.CardDAO) *api {
	return &api{
		cardRepository: repo,
	}
}

// GetCard gets a card.
func (a *api) GetCard(ctx context.Context, req *cardv1.GetCardRequest) (*cardv1.Card, error) {

	// get: fetch the card
	// TODO: verify there wasn't a not found error
	res, err := a.cardRepository.GetCard(ctx, req.GetName())
	if err != nil {
		if _, ok := err.(*cerrors.APIError); ok {
			return nil, err
		}
		return nil, &cerrors.APIError{
			Code:      codes.Internal,
			Message:   "failed to get card",
			BaseError: err,
		}
	}

	// get: return the card
	return marshalCard(res), nil

}

// ListCards gets a collection of cards.
func (a *api) ListCards(ctx context.Context, _ *cardv1.ListCardsRequest) (*cardv1.ListCardsResponse, error) {

	// list: fetch the cards
	res, err := a.cardRepository.ListCards(ctx)
	if err != nil {
		return nil, &cerrors.APIError{
			Code:      codes.Internal,
			Message:   "failed to list cards",
			BaseError: err,
		}
	}

	// list: return the card
	return &cardv1.ListCardsResponse{
		Cards: marshalCards(res),
	}, nil

}

// marshalCard translates the DAO model to an API model.
func marshalCard(card models.Card) *cardv1.Card {
	return &cardv1.Card{
		Name: card.Name(),
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
