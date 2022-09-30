package repo

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/gocommon/pkg/restclient"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/blackhorseya/todo-app/test/testdata"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type SuiteHTTP struct {
	suite.Suite
	logger     *zap.Logger
	restclient *restclient.MockRestClient
	repo       IRepo
}

func (s *SuiteHTTP) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.restclient = new(restclient.MockRestClient)

	repo, err := CreateHTTP(&Options{BaseURL: "http://localhost:8080"}, s.restclient)
	if err != nil {
		panic(err)
	}
	s.repo = repo
}

func TestSuiteHTTP(t *testing.T) {
	suite.Run(t, new(SuiteHTTP))
}

func (s *SuiteHTTP) Test_rest_GetByID() {
	type args struct {
		id   primitive.ObjectID
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *todo.Task
		wantErr  bool
	}{
		{
			name: "do request then error",
			args: args{id: testdata.TaskOID1, mock: func() {
				s.restclient.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "resp code not 200 then error",
			args: args{id: testdata.TaskOID1, mock: func() {
				data, _ := json.Marshal(response.Response{Code: 400, Msg: "failed", Data: nil})
				body := io.NopCloser(strings.NewReader(string(data)))
				s.restclient.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "resp data then success",
			args: args{id: testdata.TaskOID1, mock: func() {
				data, _ := json.Marshal(response.OK.WithData(testdata.Task1))
				body := io.NopCloser(strings.NewReader(string(data)))
				s.restclient.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil).Once()
			}},
			wantTask: testdata.Task1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTask, err := s.repo.GetByID(contextx.BackgroundWithLogger(s.logger), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("GetByID() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			s.restclient.AssertExpectations(t)
		})
	}
}
