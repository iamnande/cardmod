package limitbreakv1

import (
	"context"

	"google.golang.org/grpc/codes"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/models"
	"github.com/iamnande/cardmod/pkg/api/limitbreakv1"
)

// api is the concrete implementation of the LimitBreakAPI gRPC service.
type api struct {
	limitbreakv1.UnimplementedLimitBreakAPIServer
	limitbreakRepository daos.LimitBreakDAO
}

// New initializes the v1 LimitBreakAPI.
func New(repo daos.LimitBreakDAO) *api {
	return &api{
		limitbreakRepository: repo,
	}
}

// GetLimitBreak gets a limitbreak.
func (a *api) GetLimitBreak(ctx context.Context, req *limitbreakv1.GetLimitBreakRequest) (*limitbreakv1.LimitBreak, error) {

	// get: fetch the limitbreak
	res, err := a.limitbreakRepository.GetLimitBreak(ctx, req.GetName())
	if err != nil {
		if _, ok := err.(*cerrors.APIError); ok {
			return nil, err
		}
		return nil, &cerrors.APIError{
			Code:      codes.Internal,
			Message:   "failed to get limitbreak",
			BaseError: err,
		}
	}

	// get: return the limitbreak
	return marshalLimitBreak(res), nil

}

// ListLimitBreaks gets a collection of limitbreaks.
func (a *api) ListLimitBreaks(ctx context.Context,
	_ *limitbreakv1.ListLimitBreaksRequest) (*limitbreakv1.ListLimitBreaksResponse, error) {

	// list: fetch the limitbreaks
	res, err := a.limitbreakRepository.ListLimitBreaks(ctx)
	if err != nil {
		return nil, &cerrors.APIError{
			Code:      codes.Internal,
			Message:   "failed to list limitbreaks",
			BaseError: err,
		}
	}

	// list: return the limitbreak
	return &limitbreakv1.ListLimitBreaksResponse{
		Limitbreaks: marshalLimitBreaks(res),
	}, nil

}

// marshalLimitBreak translates the DAO model to an API model.
func marshalLimitBreak(limitbreak models.LimitBreak) *limitbreakv1.LimitBreak {
	return &limitbreakv1.LimitBreak{
		Name: limitbreak.Name(),
	}
}

// marshalLimitBreaks translates a collection of DAO models into a collection API model.
func marshalLimitBreaks(limitbreaks []models.LimitBreak) []*limitbreakv1.LimitBreak {
	res := make([]*limitbreakv1.LimitBreak, len(limitbreaks))
	for i := 0; i < len(limitbreaks); i++ {
		res[i] = marshalLimitBreak(limitbreaks[i])
	}
	return res
}
