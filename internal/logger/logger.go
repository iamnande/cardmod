package logger

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// service is the service object to log.
type service struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// NewLogger constructs a new logr compatible instance of a logger.
func NewLogger(name, version string) logr.Logger {

	// new: initialize zap config
	cfg := zapcore.EncoderConfig{
		LevelKey:     "level",
		CallerKey:    "caller",
		MessageKey:   "message",
		TimeKey:      "timestamp",
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
	}

	// new: initialize zap instance
	zapInstance := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg),
			zapcore.Lock(os.Stdout),
			zap.DebugLevel,
		),
		zap.AddCaller(),
	)

	// new: return zapr (logr) logger
	return zapr.NewLogger(zapInstance).
		WithName(name).
		WithValues("service", &service{
			Name:    name,
			Version: version,
		})

}
