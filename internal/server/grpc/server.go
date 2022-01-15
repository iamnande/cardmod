package grpc

import (
	"net"

	"github.com/iamnande/cardmod/internal/grpc/cardapi"
	"github.com/iamnande/cardmod/internal/grpc/magicapi"
	"github.com/iamnande/cardmod/pkg/api/cardv1"
	"github.com/iamnande/cardmod/pkg/api/magicv1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Server is the internal (gRPC) Server router.
type Server struct {
	version string
	port    string
	logger  *zap.Logger
	server  *grpc.Server
}

// ServerConfig is the server configuration.
type ServerConfig struct {
	Port    string
	Logger  *zap.Logger
	Version string
}

// NewServer creates a new, pre-configured, gRPC server.
func NewServer(cfg *ServerConfig) *Server {

	// initialize a new server instance
	return &Server{
		port:    cfg.Port,
		version: cfg.Version,
	}

}

// Server starts the gRPC server.
func (s *Server) Serve() error {

	// initialize gRPC server
	server := grpc.NewServer()

	// initialize listener
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}

	// add the services to the server
	cardService := cardapi.New()
	s.logger.Info("registering the CardAPI service")
	cardv1.RegisterCardAPIServer(server, &cardService)

	magicService := magicapi.New()
	s.logger.Info("registering the MagicAPI service")
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
