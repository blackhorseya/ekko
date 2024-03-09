package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config defines the config for logging.
type Config struct {
	// Level is the log level. options: debug, info, warn, error, dpanic, panic, fatal (default: info)
	Level string `json:"level" yaml:"level"`

	// Format is the log format. options: json, console (default: console)
	Format string `json:"format" yaml:"format"`
}

// InitWithConfig will initialize the logger with config.
func InitWithConfig(config Config) error {
	level := zap.NewAtomicLevel()
	err := level.UnmarshalText([]byte(config.Level))
	if err != nil {
		return err
	}

	cw := zapcore.Lock(os.Stdout)
	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	enc := zapcore.NewConsoleEncoder(zapConfig)
	if config.Format == "json" {
		zapConfig = zap.NewProductionEncoderConfig()
		enc = zapcore.NewJSONEncoder(zapConfig)
	}

	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(enc, cw, level))

	core := zapcore.NewTee(cores...)
	logger := zap.New(core)

	zap.ReplaceGlobals(logger)

	return nil
}
