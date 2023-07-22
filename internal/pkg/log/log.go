package log

import (
	"os"

	configx "github.com/blackhorseya/ekko/internal/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger serve caller to create zap.Logger
func NewLogger(cfg *configx.Config) (*zap.Logger, error) {
	var (
		err    error
		level  = zap.NewAtomicLevel()
		logger *zap.Logger
	)

	err = level.UnmarshalText([]byte(cfg.Log.Level))
	if err != nil {
		return nil, err
	}

	cw := zapcore.Lock(os.Stdout)
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	enc := zapcore.NewConsoleEncoder(config)
	if cfg.Log.Format == "json" {
		config = zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.RFC3339NanoTimeEncoder
		enc = zapcore.NewJSONEncoder(config)
	}

	cores := make([]zapcore.Core, 0, 2)
	cores = append(cores, zapcore.NewCore(enc, cw, level))

	core := zapcore.NewTee(cores...)
	logger = zap.New(core)

	zap.ReplaceGlobals(logger)

	return logger, nil
}
