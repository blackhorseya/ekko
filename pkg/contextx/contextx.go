package contextx

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Key declare key string
type Key string

var (
	// KeyCtx ctx key string
	KeyCtx = Key("ctx")
)

// Contextx extends google's context to support logging methods
type Contextx struct {
	context.Context
	logger *zap.Logger
}

// Background returns a non-nil, empty Context. It is never canceled, has no values, and
// has no deadline. It is typically used by the main function, initialization, and tests,
// and as the top-level Context for incoming requests
func Background() Contextx {
	return Contextx{
		Context: context.Background(),
		logger:  nil,
	}
}

// BackgroundWithLogger returns a non-nil, empty Context. It is never canceled, has no values, and
// has no deadline. It is typically used by the main function, initialization, and tests,
// and as the top-level Context for incoming requests
func BackgroundWithLogger(logger *zap.Logger) Contextx {
	return Contextx{
		Context: context.Background(),
		logger:  logger,
	}
}

// WithValue returns a copy of parent in which the value associated with key is val.
func WithValue(parent Contextx, key interface{}, val interface{}) Contextx {
	return Contextx{
		Context: context.WithValue(parent, key, val),
	}
}

// WithCancel returns a copy of parent with added cancel function
func WithCancel(parent Contextx) (Contextx, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)
	return Contextx{
		Context: ctx,
	}, cancel
}

// WithTimeout returns a copy of parent with timeout condition and cancel function
func WithTimeout(parent Contextx, d time.Duration) (Contextx, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(parent, d)
	return Contextx{
		Context: ctx,
	}, cancel
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (c *Contextx) Debug(msg string, fields ...zapcore.Field) {
	if c.logger == nil {
		return
	}

	c.logger.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (c *Contextx) Info(msg string, fields ...zapcore.Field) {
	if c.logger == nil {
		return
	}

	c.logger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (c *Contextx) Warn(msg string, fields ...zapcore.Field) {
	if c.logger == nil {
		return
	}

	c.logger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (c *Contextx) Error(msg string, fields ...zapcore.Field) {
	if c.logger == nil {
		return
	}

	c.logger.Error(msg, fields...)
}

// Elapsed calculate method execution time
func (c *Contextx) Elapsed(msg string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		if c.logger == nil {
			return
		}

		c.logger.Debug(msg, zap.Duration("elapsed", elapsed))
	}
}
