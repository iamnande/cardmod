package server

import (
	"net"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	calculatev1api "github.com/iamnande/cardmod/internal/api/calculatev1"
	cardv1api "github.com/iamnande/cardmod/internal/api/cardv1"
	itemv1api "github.com/iamnande/cardmod/internal/api/itemv1"
	limitbreakv1api "github.com/iamnande/cardmod/internal/api/limitbreakv1"
	magicv1api "github.com/iamnande/cardmod/internal/api/magicv1"
	refinementv1api "github.com/iamnande/cardmod/internal/api/refinementv1"
	"github.com/iamnande/cardmod/internal/config"
	"github.com/iamnande/cardmod/pkg/api/calculatev1"
	"github.com/iamnande/cardmod/pkg/api/cardv1"
	"github.com/iamnande/cardmod/pkg/api/itemv1"
	"github.com/iamnande/cardmod/pkg/api/limitbreakv1"
	"github.com/iamnande/cardmod/pkg/api/magicv1"
	"github.com/iamnande/cardmod/pkg/api/refinementv1"
)

// Config is the server configuration.
type Config struct {
	Config  *config.Config
	Logger  logr.Logger
	Version string
}

// Server is the server instance.
type Server struct {

	// configurable attributes
	config  *config.Config
	logger  logr.Logger
	version string

	// constructed attributes
	server *grpc.Server
}

// NewServer initializes a new gRPC server instance using the provided configuration.
func NewServer(cfg *Config) *Server {
	return &Server{
		config:  cfg.Config,
		logger:  cfg.Logger,
		version: cfg.Version,
	}
}

// Server will start the server listener.
func (s *Server) Serve() error {

	// serve: construct underlying gRPC server instance
	srvr := grpc.NewServer(grpc.ChainUnaryInterceptor(
		s.LoggingInterceptor(),
	))

	// serve: register APIs
	cardv1.RegisterCardAPIServer(srvr, cardv1api.New())
	itemv1.RegisterItemAPIServer(srvr, itemv1api.New())
	limitbreakv1.RegisterLimitBreakAPIServer(srvr, limitbreakv1api.New())
	magicv1.RegisterMagicAPIServer(srvr, magicv1api.New())
	refinementv1.RegisterRefinementAPIServer(srvr, refinementv1api.New())
	calculatev1.RegisterCalculateAPIServer(srvr, calculatev1api.New())

	// serve: enable reflection
	reflection.Register(srvr)

	// serve: construct the listener
	lis, err := net.Listen("tcp", s.config.Server.PortListener())
	if err != nil {
		return err
	}

	// serve: run the server
	if err := srvr.Serve(lis); err != nil {
		return err
	}

	// serve: assign and return
	s.server = srvr
	return nil

}

// GracefullyStop attempts to gracefully stop the underlying server.
func (s *Server) GracefullyStop() {
	if s.server != nil {
		s.server.GracefulStop()
		s.server = nil
	}
}

// Stop will forcefully stop the underlying server.
func (s *Server) Stop() {
	if s.server != nil {
		s.server.Stop()
		s.server = nil
	}
}
