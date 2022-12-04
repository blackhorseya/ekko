package log

import (
	"os"

	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Options declare log's configuration
type Options struct {
	Level    string
	Encoding string
}

// NewOptions serve caller to create Options
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("log", o); err != nil {
		return nil, err
	}

	return o, nil
}

// NewLogger serve caller to create zap.Logger
func NewLogger(o *Options) (*zap.Logger, error) {
	var (
		err    error
		level  = zap.NewAtomicLevel()
		logger *zap.Logger
	)

	err = level.UnmarshalText([]byte(o.Level))
	if err != nil {
		return nil, err
	}

	cw := zapcore.Lock(os.Stdout)
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	enc := zapcore.NewConsoleEncoder(config)
	if o.Encoding == "json" {
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

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewLogger, NewOptions)
