package config

import (
	"fmt"
	"time"
)

const (

	// appEnvPrefix is the application level environment variable prefix to search under.
	appEnvPrefix = "CARDMOD_"

	// defaultEnvironment is the default environment name.
	defaultEnvironment = "local"

	// defaultServerPort is the default port listener for the gRPC server.
	defaultServerPort = 9000

	// defaultServerShutdownGracePeriod is the default timeout of the graceful shutdown.
	defaultServerShutdownGracePeriod = time.Second * 5
)

// Config is the configuration mechanism used for various instantiations
// of various components within the service ecosystem.
type Config struct {

	// Environment is the name of the environment where the service is running.
	Environment string `env:"ENVIRONMENT" json:"environment" yaml:"environment"`

	// Server is the server specific configurations.
	Server *ServerConfig `envPrefix:"SERVER_" json:"server" yaml:"server"`
}

// ServerConfig is the gRPC server configuration.
type ServerConfig struct {

	// Port is the port the gRPC server will listen on.
	Port int `env:"PORT" json:"port" yaml:"port"`

	// ShutdownGracePeriod is the grace period, in seconds, before the server will
	// be forcefully shut down.
	ShutdownGracePeriod time.Duration `env:"SHUTDOWN_GRACE_PERIOD" json:"shutdown_grace_period" yaml:"shutdown_grace_period"`
}

// PortListener is a stringer method on the Port attribute.
func (sc *ServerConfig) PortListener() string {
	return fmt.Sprintf(":%d", sc.Port)
}
