package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/todo-app/internal/pkg/infra/transports/http"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Application declare application's information
type Application struct {
	name       string
	logger     *zap.Logger
	httpServer *http.Server
}

// Option declare application options
type Option func(app *Application) error

// HTTPServerOption declare http server option
func HTTPServerOption(svr *http.Server) Option {
	return func(app *Application) error {
		svr.Application(app.name)
		app.httpServer = svr

		return nil
	}
}

// New serve caller to new Application
func New(name string, logger *zap.Logger, options ...Option) (*Application, error) {
	app := &Application{
		name:   name,
		logger: logger.With(zap.String("type", "Application")),
	}

	for _, option := range options {
		if err := option(app); err != nil {
			return nil, err
		}
	}

	return app, nil
}

// Start serve caller to start an application
func (a *Application) Start() error {
	if a.httpServer != nil {
		if err := a.httpServer.Start(); err != nil {
			return errors.Wrap(err, "http server start error")
		}
	}

	return nil
}

// AwaitSignal serve caller to await server running
func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		a.logger.Info("receive a signal", zap.String("signal", s.String()))
		if a.httpServer != nil {
			if err := a.httpServer.Stop(); err != nil {
				a.logger.Warn("stop http server error", zap.Error(err))
			}
		}

		os.Exit(0)
	}
}
