package main

import (
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/pkg/errors"
)

type Service struct {
	cli adapters.CLI
}

// NewService serve caller to create service instance
func NewService(cli adapters.CLI) (*Service, error) {
	svc := &Service{
		cli: cli,
	}

	return svc, nil
}

func (s *Service) Start() error {
	if s.cli == nil {
		return errors.New("cli instance is nil")
	}

	return s.cli.Execute()
}
