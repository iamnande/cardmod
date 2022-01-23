package grpc

import (
	"net"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"

	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/grpc/cardapi"
	"github.com/iamnande/cardmod/internal/grpc/livezapi"
	"github.com/iamnande/cardmod/internal/grpc/magicapi"
	"github.com/iamnande/cardmod/pkg/api/cardv1"
	"github.com/iamnande/cardmod/pkg/api/livezv1"
	"github.com/iamnande/cardmod/pkg/api/magicv1"
)

// Server is the internal (gRPC) Server router.
type Server struct {
	version string
	port    string
	logger  logr.Logger
	server  *grpc.Server

	// repositories
	cardRepository  daos.CardDAO
	magicRepository daos.MagicDAO
}

// ServerConfig is the server configuration.
type ServerConfig struct {
	Port    string
	Logger  logr.Logger
	Version string

	// Repositories
	CardRepository  daos.CardDAO
	MagicRepository daos.MagicDAO
}

// NewServer creates a new, pre-configured, gRPC server.
func NewServer(cfg *ServerConfig) *Server {

	// initialize a new server instance
	return &Server{
		port:    cfg.Port,
		logger:  cfg.Logger,
		version: cfg.Version,

		// repositories
		cardRepository:  cfg.CardRepository,
		magicRepository: cfg.MagicRepository,
	}

}

// Server starts the gRPC server.
func (s *Server) Serve() error {

	// initialize gRPC server
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		s.LoggingInterceptor(),
	))

	// initialize listener
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}

	// add the services to the server
	livezService := livezapi.New()
	livezv1.RegisterLivezAPIServer(server, &livezService)

	cardService := cardapi.New(s.cardRepository)
	cardv1.RegisterCardAPIServer(server, &cardService)

	magicService := magicapi.New(s.magicRepository)
	magicv1.RegisterMagicAPIServer(server, &magicService)

	// start the server
	s.server = server
	return server.Serve(lis)

}

// Stop will stop the gRPC gracefully.
func (s *Server) Stop() {
	s.server.GracefulStop()
	s.server = nil
}
