package repo

import (
	"fmt"
	"testing"
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type suiteRepo struct {
	suite.Suite
	pool     *dockertest.Pool
	resource *dockertest.Resource
	client   *mongo.Client
	repo     IRepo
}

func (s *suiteRepo) SetupTest() {
	logger := zap.NewNop()
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}
	s.pool = pool

	resource, err := pool.Run("mongo", "4.4.10", nil)
	if err != nil {
		panic(err)
	}
	s.resource = resource

	err = pool.Retry(func() error {
		uri := fmt.Sprintf("mongodb://localhost:%s/", resource.GetPort("27017/tcp"))
		s.client, err = mongo.Connect(contextx.Background(), options.Client().ApplyURI(uri))
		if err != nil {
			return err
		}

		return s.client.Ping(contextx.Background(), readpref.Primary())
	})
	if err != nil {
		panic(err)
	}

	repo, err := CreateIRepo(logger, s.client)
	if err != nil {
		panic(err)
	}
	s.repo = repo
}

func (s *suiteRepo) TearDownTest() {
	_ = s.client.Disconnect(contextx.Background())
	_ = s.pool.Purge(s.resource)
}

func TestSuiteRepo(t *testing.T) {
	suite.Run(t, new(suiteRepo))
}

func (s *suiteRepo) Test_impl_Ping() {
	type args struct {
		timeout time.Duration
		mock    func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "ping then success",
			args:    args{timeout: 5 * time.Second},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.Ping(contextx.Background(), tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
