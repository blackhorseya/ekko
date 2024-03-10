//go:build external

package mongodb

import (
	"testing"

	"github.com/blackhorseya/ekko/entity/domain/workflow/repo"
	"github.com/stretchr/testify/suite"
)

type suiteExternal struct {
	suite.Suite

	repo repo.IIssueRepo
}

func TestExternal(t *testing.T) {
	suite.Run(t, new(suiteExternal))
}
