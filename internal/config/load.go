package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
)

var (
	// config is the private instance loaded during API instantiation.
	config *Config

	// once is used to ensure the configuration is loaded only one time.
	once sync.Once
)

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

		// core
		Environment: defaultEnvironment,

		// server
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
