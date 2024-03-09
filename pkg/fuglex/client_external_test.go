//go:build external

package fuglex

import (
	"testing"

	"github.com/blackhorseya/ekko/pkg/configx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteClientExternal struct {
	suite.Suite

	client Dialer
}

func (s *suiteClientExternal) SetupTest() {
	err := configx.LoadWithPathAndName("", "orianna")
	s.Require().NoError(err)

	client, err := NewClient()
	s.Require().NoError(err)
	s.client = client
}

func TestClientExternal(t *testing.T) {
	suite.Run(t, new(suiteClientExternal))
}

func (s *suiteClientExternal) TestClient_IntradayQuote() {
	ctx := contextx.Background()
	res, err := s.client.IntradayQuote(ctx, "2330")
	s.Require().NoError(err)
	s.Require().NotNil(res)
	ctx.Debug("print got", zap.Any("got", &res))
}
