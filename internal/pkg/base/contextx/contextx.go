package contextx

import (
	"context"
	"time"
)

// Contextx extends google's context to support logging methods
type Contextx struct {
	context.Context
}

// Background returns a non-nil, empty Context. It is never canceled, has no values, and
// has no deadline. It is typically used by the main function, initialization, and tests,
// and as the top-level Context for incoming requests
func Background() Contextx {
	return Contextx{
		Context: context.Background(),
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
