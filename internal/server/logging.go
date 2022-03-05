package server

import (
	"context"
	"path"
	"time"

	"github.com/go-logr/logr"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/iamnande/cardmod/internal/cerrors"
)

// request is the request data to log.
// TODO: request ID middleware and inject it here.
type request struct {
	Method  string `json:"method"`
	Service string `json:"service"`
}

// response is the response data to log.
type response struct {
	Code     string        `json:"code"`
	Duration time.Duration `json:"duration"`
	Error    error         `json:"error"`
}

// LoggingInterceptor is the gRPC server access logging interceptor (middleware).
func (s *Server) LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {

		// log: extract request data
		start := time.Now().UTC()
		method := path.Base(info.FullMethod)
		service := path.Dir(info.FullMethod)[1:]

		// log: inject the request logger into the context
		log := s.logger.WithValues("request", &request{
			Method:  method,
			Service: service,
		})
		reqCTX := logr.NewContext(ctx, log)

		// log: execute handler
		var apiError *cerrors.APIError
		res, err = handler(reqCTX, req)
		if err != nil {
			switch apiErr := err.(type) {
			case *cerrors.APIError:
				apiError = apiErr
				err = status.Error(apiErr.Code, apiErr.Message)
			}
		}

		// log: construct the response log
		resLog := &response{
			Code:     grpc_logging.DefaultErrorToCode(err).String(),
			Duration: time.Since(start),
			Error:    err,
		}
		if apiError != nil {
			resLog.Error = apiError
		}
		log = log.WithValues("response", resLog)

		// log: log the request/response
		if err != nil &&
			resLog.Code != codes.InvalidArgument.String() &&
			resLog.Code != codes.AlreadyExists.String() &&
			resLog.Code != codes.NotFound.String() {
			log.Error(nil, "error")
		} else {
			log.Info("success")
		}

		// log: close the request cycle
		return res, err

	}
}
