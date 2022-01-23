package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
)

const (
	// DefaultGRPCPort is the default port listener for the gRPC server.
	DefaultGRPCPort = ":9000"

	// DefaultEnvironment is the default environment name.
	DefaultEnvironment = "local"

	// DefaultDatabaseEndpoint is the default database connection endpoint string.
	DefaultDatabaseEndpoint = "postgres://localhost:5432/cardmod?sslmode=disable&user=postgres&password=postgres"
)

var (
	// config is the private instance loaded during API instantiation.
	config *Config

	// once is used to ensure the configuration is loaded only one time.
	once sync.Once
)

// Config is a protected configuration mechanism used for various instantiations
// of various components within the API ecosystem.
type Config struct {

	// GRPCPort is the port the gRPC server will listen on.
	GRPCPort string `env:"GRPC_PORT"`

	// Environment is the name of the environment where the service is running.
	Environment string `env:"ENVIRONMENT"`

	// DatabaseEndpoint is the database endpoint used for connection. This is
	// currently configured to be a DSN (relational database speak for
	// connection details).
	DatabaseEndpoint string `env:"DATABASE_ENDPOINT"`
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
		GRPCPort:         DefaultGRPCPort,
		Environment:      DefaultEnvironment,
		DatabaseEndpoint: DefaultDatabaseEndpoint,
	}

	// load: extract attributes from environment variables
	// NOTE: unless there are required attributes that are not supplied, or
	//       there are explicit non-string types, this library will not error.
	if err := env.Parse(cfg, env.Options{
		Prefix: "CARDMOD_",
	}); err != nil {
		panic(err)
	}

	// load: return constructed Config instance
	return cfg

}
