package grpc

import (
	"context"
	"path"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type request struct {
	RequestID string `json:"id"`
	Method    string `json:"method"`
	Service   string `json:"service"`
}

type response struct {
	Code     string      `json:"code"`
	Duration string      `json:"duration"`
	Errors   interface{} `json:"errors,omitempty"`
}

func (s *Server) LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		// log: extract request data
		start := time.Now().UTC()
		method := path.Base(info.FullMethod)
		service := path.Dir(info.FullMethod)[1:]

		// log: create request-specific logger
		log := s.logger.WithValues("request", &request{
			Method:    method,
			Service:   service,
			RequestID: uuid.NewString(),
		})
		newCTX := logr.NewContext(ctx, log)

		// log: execute handler
		res, err := handler(newCTX, req)

		// log: extract response information
		stat, _ := status.FromError(err)
		code := grpc_logging.DefaultErrorToCode(err)
		resLog := &response{
			Code:     code.String(),
			Duration: time.Since(start).String(),
		}
		if stat.Details() != nil {
			resLog.Errors = stat.Details()
		}
		log = log.WithValues("response", resLog)

		// log: write either an error or an info log
		if err != nil {
			// note: we're passing nil here because we want the error string to be
			// under the response object (using stat.Message()) while still preserving
			// the ERROR leveling in the log message.
			log.Error(nil, "error")
		} else {
			log.Info("success")
		}

		// log: close the request cycle
		return res, err

	}
}
