package repo

import (
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteHTTPClient struct {
	suite.Suite
	logger *zap.Logger
	repo   IRepo
}
