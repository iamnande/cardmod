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

	// defaultDatabaseHostname is the default database hostname.
	defaultDatabaseHostname = "localhost"

	// defaultDatabasePort is the default database port.
	defaultDatabasePort = 5432

	// defaultDatabaseName is the default database name.
	defaultDatabaseName = "cardmod"

	// defaultDatabaseSSLMode is the default database SSL mode.
	defaultDatabaseSSLMode = "disable"

	// defaultDatabaseUsername is the default database username.
	defaultDatabaseUsername = "cardmod"

	// defaultDatabasePassword is the default database password.
	defaultDatabasePassword = "cardmod"

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

	// Database is the database specific configurations.
	Database *DatabaseConfig `envPrefix:"DATABASE_" json:"database" yaml:"database"`

	// Server is the server specific configurations.
	Server *ServerConfig `envPrefix:"SERVER_" json:"server" yaml:"server"`
}

// DatabaseConfig is the database configuration.
type DatabaseConfig struct {

	// Hostname is the hostname of the database.
	Hostname string `env:"HOSTNAME" json:"hostname" yaml:"hostname"`

	// Port is the port of the database.
	Port int `env:"PORT" json:"port" yaml:"port"`

	// Name is the name of the database to connect to.
	Name string `env:"NAME" json:"name" yaml:"name"`

	// SSLMode is the database SSL mode used for connection.
	SSLMode string `env:"SSL_MODE" json:"ssl_mode" yaml:"ssl_mode"`

	// Username is the username to connect to the database with.
	Username string `env:"USERNAME" json:"username" yaml:"username"`

	// Password is the password to connect to the database with.
	Password string `env:"PASSWORD" json:"password" yaml:"password"`
}

// DSN is a stringer method on the DatabaseConfig attributes to construct a DSN.
func (dc *DatabaseConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%d/%s?sslmode=%s&user=%s&password=%s",
		dc.Hostname, dc.Port, dc.Name, dc.SSLMode, dc.Username, dc.Password)
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
