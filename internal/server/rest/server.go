package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-logr/logr"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/iamnande/cardmod/pkg/api/calculationv1"
	"github.com/iamnande/cardmod/pkg/api/cardv1"
	"github.com/iamnande/cardmod/pkg/api/healthv1"
	"github.com/iamnande/cardmod/pkg/api/magicv1"
)

// Server is the internal REST API Server.
type Server struct {
	ctx          context.Context
	logger       logr.Logger
	grpcEndpoint string
	server       http.Server
}

// ServerConfig is the configuration mechanism for the *server.
type ServerConfig struct {
	Context      context.Context
	Logger       logr.Logger
	GRPCEndpoint string
	RESTEndpoint string
}

// NewServer initializes a new REST API server using the configuration provided.
func NewServer(cfg *ServerConfig) (*Server, error) {

	// server: initialize REST router
	router := chi.NewRouter()

	// server: configure the gRPC client (HTTP->gRPC)
	// TODO: need to figure out gRPC failure to REST failure
	// TODO: need to figure out REST error handling/model
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	connection, err := grpc.Dial(cfg.GRPCEndpoint, opts...)
	if err != nil {
		return nil, err
	}

	// server: register the gRPC gateway handler(s)
	// TODO: wrap this in a loop or something
	if err := healthv1.RegisterHealthAPIHandler(cfg.Context, mux, connection); err != nil {
		return nil, err
	}
	if err := cardv1.RegisterCardAPIHandler(cfg.Context, mux, connection); err != nil {
		return nil, err
	}
	if err := magicv1.RegisterMagicAPIHandler(cfg.Context, mux, connection); err != nil {
		return nil, err
	}
	if err := calculationv1.RegisterCalculationAPIHandler(cfg.Context, mux, connection); err != nil {
		return nil, err
	}

	// server: mount the gRPC gateway into the REST server
	router.Mount("/v1", mux)

	// server: return pre-configured REST server (with gRPC translation)
	return &Server{
		ctx:          cfg.Context,
		logger:       cfg.Logger,
		grpcEndpoint: cfg.GRPCEndpoint,
		server: http.Server{
			Addr:    cfg.RESTEndpoint,
			Handler: router,
		},
	}, nil

}

// Server starts the REST server.
func (s *Server) Serve() error {
	return s.server.ListenAndServe()
}

// Stop will stop the gRPC gracefully.
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}