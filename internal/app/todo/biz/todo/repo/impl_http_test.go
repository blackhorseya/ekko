package repo

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/gocommon/pkg/restclient"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/ticket"
	"github.com/blackhorseya/todo-app/test/testdata"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteHTTP struct {
	suite.Suite
	logger     *zap.Logger
	restclient *restclient.MockRestClient
	repo       ITodoRepo
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
		id   uint64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *ticket.Task
		wantErr  bool
	}{
		{
			name: "do request then error",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.restclient.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "resp code not 200 then error",
			args: args{id: testdata.Task1.ID, mock: func() {
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
			args: args{id: testdata.Task1.ID, mock: func() {
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

func (s *SuiteHTTP) Test_rest_List() {
	type args struct {
		condition QueryTodoCondition
		mock      func()
	}
	tests := []struct {
		name      string
		args      args
		wantTasks []*ticket.Task
		wantErr   bool
	}{
		{
			name: "do request then error",
			args: args{condition: QueryTodoCondition{Limit: 10, Offset: 0}, mock: func() {
				s.restclient.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantTasks: nil,
			wantErr:   true,
		},
		{
			name: "resp code not 200 then error",
			args: args{condition: QueryTodoCondition{Limit: 10, Offset: 0}, mock: func() {
				data, _ := json.Marshal(response.Response{Code: 400, Msg: "failed", Data: nil})
				body := io.NopCloser(strings.NewReader(string(data)))
				s.restclient.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil).Once()
			}},
			wantTasks: nil,
			wantErr:   true,
		},
		{
			name: "list tasks then success",
			args: args{condition: QueryTodoCondition{Limit: 10, Offset: 0}, mock: func() {
				data, _ := json.Marshal(response.Response{Code: 200, Msg: "ok", Data: []*ticket.Task{testdata.Task1}})
				body := io.NopCloser(strings.NewReader(string(data)))
				s.restclient.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil).Once()
			}},
			wantTasks: []*ticket.Task{testdata.Task1},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTasks, err := s.repo.List(contextx.BackgroundWithLogger(s.logger), tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTasks, tt.wantTasks) {
				t.Errorf("List() gotTasks = %v, want %v", gotTasks, tt.wantTasks)
			}

			s.restclient.AssertExpectations(t)
		})
	}
}

func (s *SuiteHTTP) Test_rest_Create() {
	type args struct {
		newTask *ticket.Task
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *ticket.Task
		wantErr  bool
	}{
		{
			name: "do request then error",
			args: args{newTask: testdata.Task1, mock: func() {
				s.restclient.On("Do", mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "resp code not 200 then error",
			args: args{newTask: testdata.Task1, mock: func() {
				data, _ := json.Marshal(response.Response{Code: 50010, Msg: "error"})
				body := io.NopCloser(bytes.NewReader(data))
				s.restclient.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 200,
					Body:       body,
				}, nil).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "create then success",
			args: args{newTask: testdata.Task1, mock: func() {
				data, _ := json.Marshal(response.Response{Code: 200, Msg: "ok", Data: testdata.Task1})
				body := io.NopCloser(bytes.NewReader(data))
				s.restclient.On("Do", mock.Anything).Return(&http.Response{
					StatusCode: 201,
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

			gotTask, err := s.repo.Create(contextx.BackgroundWithLogger(s.logger), tt.args.newTask)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("Create() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			s.restclient.AssertExpectations(t)
		})
	}
}
