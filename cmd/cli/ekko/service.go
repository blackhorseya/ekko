package main

import (
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	cli    adapters.CLI
}

// NewService serve caller to create service instance
func NewService(logger *zap.Logger, cli adapters.CLI) (*Service, error) {
	svc := &Service{
		logger: logger.With(zap.String("type", "service")),
		cli:    cli,
	}

	return svc, nil
}

func (s *Service) Start() error {
	if s.cli == nil {
		return errors.New("cli instance is nil")
	}

	return s.cli.Execute()
}
