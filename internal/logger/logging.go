package logger

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
)

const (
	// ServiceKey is the key of the service in the log message.
	ServiceKey = "service"

	// TimestampKey is the key of the timestamp of the log message.
	TimestampKey = "timestamp"
)

// service is the service instance using the logger.
type service struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// NewLogger creates a new logr instance. Under the hood it's using zerolog.
func NewLogger(name, version string) logr.Logger {

	// logger: initialize zerolog
	zerolog.TimestampFieldName = TimestampKey

	log := zerolog.New(os.Stdout).With().Timestamp().Stack().Logger()

	// logger: initialize zapr (logr implementation)
	return zerologr.New(&log).WithValues(ServiceKey, &service{
		Name:    name,
		Version: version,
	})

}
