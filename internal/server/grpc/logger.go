package grpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// TODO: clean up the log format
// TODO: can we clean up the callers too (aka calling handler)?
// TODO: clean up the fields
// TODO: log errors and structures.
func (s *Server) LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		// log: extract request data
		fields := make([]zap.Field, 0)
		fields = append(fields, zap.String("request.method", info.FullMethod))

		// log: execute handler
		response, err := handler(ctx, req)

		// log: extract call/response data
		s.logger.Info("request processed", fields...)

		// log: continue to next handler (if applicable)
		return response, err

	}
}
