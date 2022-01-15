package healthapi

import (
	"context"

	"github.com/iamnande/cardmod/pkg/api/healthv1"
)

// api is the service implementation of the generated LivezAPI gRPC service.
type api struct {
	healthv1.UnimplementedHealthAPIServer
}

// New initializes a new livez api instance.
func New() api {
	return api{}
}

// Check does a simple heartbeat check on the service.
func (api *api) Check(_ context.Context, _ *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	return &healthv1.HealthCheckResponse{
		Status: healthv1.ServiceStatus_SERVING,
	}, nil
}
