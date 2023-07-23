package app

import (
	"os"
	"os/signal"
	"syscall"

	http "github.com/blackhorseya/ekko/internal/pkg/httpx"
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/blackhorseya/ekko/pkg/httpx"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ProviderSet define the provider set
var ProviderSet = wire.NewSet(NewService, http.ServerSet, NewRestful)

// Service define the restful service
type Service struct {
	logger     *zap.Logger
	httpserver httpx.Server
}

// NewService will create a restful service
func NewService(logger *zap.Logger, httpserver httpx.Server, restful adapters.Restful) *Service {
	if restful != nil {
		restful.InitRouting()
	}

	return &Service{
		logger:     logger,
		httpserver: httpserver,
	}
}

// Start will start the restful service
func (s *Service) Start() error {
	s.logger.Info("starting service")

	if s.httpserver != nil {
		err := s.httpserver.Start()
		if err != nil {
			return errors.Wrap(err, "start http server")
		}
	}

	return nil
}

// AwaitSignal will await the signal
func (s *Service) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		s.logger.Info("receive signal", zap.String("signal", sig.String()))

		if s.httpserver != nil {
			err := s.httpserver.Stop()
			if err != nil {
				s.logger.Warn("stop http server", zap.Error(err))
			}
		}

		os.Exit(0)
	}

	return nil
}
