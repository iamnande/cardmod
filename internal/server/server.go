package server

import (
	"net"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	cardv1api "github.com/iamnande/cardmod/internal/apis/cardv1"
	itemv1api "github.com/iamnande/cardmod/internal/apis/itemv1"
	limitbreakv1api "github.com/iamnande/cardmod/internal/apis/limitbreakv1"
	magicv1api "github.com/iamnande/cardmod/internal/apis/magicv1"
	refinementv1api "github.com/iamnande/cardmod/internal/apis/refinementv1"
	"github.com/iamnande/cardmod/internal/config"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/repositories"
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

	DatabaseClient *database.Client
}

// Server is the server instance.
type Server struct {

	// configurable attributes
	config         *config.Config
	logger         logr.Logger
	version        string
	databaseClient *database.Client

	// constructed attributes
	server *grpc.Server
}

// NewServer initializes a new gRPC server instance using the provided configuration.
func NewServer(cfg *Config) *Server {
	return &Server{
		config:         cfg.Config,
		logger:         cfg.Logger,
		version:        cfg.Version,
		databaseClient: cfg.DatabaseClient,
	}
}

// Server will start the server listener.
func (s *Server) Serve() error {

	// serve: construct underlying gRPC server instance
	srvr := grpc.NewServer(grpc.ChainUnaryInterceptor(
		s.LoggingInterceptor(),
	))

	// serve: initialize APIs
	cardAPI := cardv1api.New(repositories.NewCardRepository(s.databaseClient))
	itemAPI := itemv1api.New(repositories.NewItemRepository(s.databaseClient))
	limitbreakAPI := limitbreakv1api.New(repositories.NewLimitBreakRepository(s.databaseClient))
	magicAPI := magicv1api.New(repositories.NewMagicRepository(s.databaseClient))
	refinementAPI := refinementv1api.New(repositories.NewRefinementRepository(s.databaseClient))

	// serve: register APIs
	cardv1.RegisterCardAPIServer(srvr, cardAPI)
	itemv1.RegisterItemAPIServer(srvr, itemAPI)
	limitbreakv1.RegisterLimitBreakAPIServer(srvr, limitbreakAPI)
	magicv1.RegisterMagicAPIServer(srvr, magicAPI)
	refinementv1.RegisterRefinementAPIServer(srvr, refinementAPI)

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
