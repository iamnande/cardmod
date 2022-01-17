package livezapi

import (
	"context"

	"github.com/iamnande/cardmod/pkg/api/livezv1"
)

// api is the service implementation of the generated LivezAPI gRPC service.
type api struct {
	livezv1.UnimplementedLivezAPIServer
}

// New initializes a new livez api instance.
func New() api {
	return api{}
}

// GetLivez does a simple heartbeat check on the service.
func (api *api) GetLivez(_ context.Context, _ *livezv1.GetLivezRequest) (*livezv1.GetLivezResponse, error) {
	return &livezv1.GetLivezResponse{
		Status: livezv1.Status_STATUS_RUNNING,
	}, nil
}
