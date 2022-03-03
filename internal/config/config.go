package config

import (
	"sync"
	"time"

	"github.com/caarlos0/env/v6"
)

const (

	// appEnvPrefix is the application level environment variable prefix to search under.
	appEnvPrefix = "CARDMOD_"

	// DefaultEnvironment is the default environment name.
	defaultEnvironment = "local"

	// DefaultServerPort is the default port listener for the gRPC server.
	defaultServerPort = 9000

	// DefaultServerShutdownGracePeriod is the default timeout of the graceful shutdown.
	defaultServerShutdownGracePeriod = time.Second * 5
)

var (
	// config is the private instance loaded during API instantiation.
	config *Config

	// once is used to ensure the configuration is loaded only one time.
	once sync.Once
)

// Config is the configuration mechanism used for various instantiations
// of various components within the service ecosystem.
type Config struct {

	// Environment is the name of the environment where the service is running.
	Environment string `env:"ENVIRONMENT" json:"environment" yaml:"environment"`

	// Server is the server specific configurations.
	Server *ServerConfig `envPrefix:"SERVER_" json:"server" yaml:"server"`
}

// ServerConfig is the gRPC server configuration mechanism.
type ServerConfig struct {

	// Port is the port the gRPC server will listen on.
	Port int `env:"PORT" json:"port" yaml:"port"`

	// ShutdownGracePeriod is the grace period, in seconds, before the server will
	// be forcefully shut down.
	ShutdownGracePeriod time.Duration `env:"SHUTDOWN_GRACE_PERIOD" json:"shutdown_grace_period" yaml:"shutdown_grace_period"`
}

// MustLoad will load the configuration from the runtime environment and
// will panic if unable to do so.
func MustLoad() *Config {
	once.Do(func() {
		config = mustLoad()
	})
	return config
}

// mustLoad is the underlying doer of fetching the *Config from the
// environment.
func mustLoad() *Config {

	// load: initialize new instance of config
	cfg := &Config{
		Environment: defaultEnvironment,
		Server: &ServerConfig{
			Port:                defaultServerPort,
			ShutdownGracePeriod: defaultServerShutdownGracePeriod,
		},
	}

	// load: extract attributes from environment variables
	// NOTE: unless there are required attributes that are not supplied, which
	// is not how we use this library, this library will not error.
	_ = env.Parse(cfg, env.Options{
		Prefix: appEnvPrefix,
	})

	// load: return constructed Config instance
	return cfg

}
