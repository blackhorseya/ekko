package repo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/blackhorseya/ekko/pkg/contextx"
	im "github.com/blackhorseya/ekko/pkg/entity/domain/issue/model"
	"github.com/blackhorseya/ekko/pkg/httpx"
	"github.com/blackhorseya/ekko/pkg/response"
	"github.com/blackhorseya/ekko/test/testdata"
	"github.com/golang/mock/gomock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteHTTPClient struct {
	suite.Suite
	ctrl   *gomock.Controller
	logger *zap.Logger

	client *httpx.MockClient
	repo   IRepo
}

func (s *suiteHTTPClient) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.ctrl = gomock.NewController(s.T())

	s.client = httpx.NewMockClient(s.ctrl)
	v := viper.GetViper()
	v.Set("app.url", "http://localhost:8080/api")
	s.repo, _ = NewHTTPClient(v, s.client)
}

func (s *suiteHTTPClient) TearDownTest() {
	s.ctrl.Finish()
}

func TestHTTPClient(t *testing.T) {
	suite.Run(t, new(suiteHTTPClient))
}

func (s *suiteHTTPClient) Test_httpclient_GetByID() {
	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *im.Ticket
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{id: testdata.Ticket1.Id, mock: func() {
				req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/tasks/%v", testdata.Ticket1.Id), nil)
				s.client.EXPECT().Do(req).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "not found then nil",
			args: args{id: testdata.Ticket1.Id, mock: func() {
				req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/tasks/%v", testdata.Ticket1.Id), nil)
				payload, _ := json.Marshal(response.Err.WithMessage("issue not found"))
				body := io.NopCloser(bytes.NewReader(payload))
				s.client.EXPECT().Do(req).Return(&http.Response{
					StatusCode: http.StatusNotFound,
					Body:       body,
				}, nil).Times(1)
			}},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "ok",
			args: args{id: testdata.Ticket1.Id, mock: func() {
				req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/tasks/%v", testdata.Ticket1.Id), nil)
				payload, _ := json.Marshal(response.OK.WithData(testdata.Ticket1))
				body := io.NopCloser(bytes.NewReader(payload))
				s.client.EXPECT().Do(req).Return(&http.Response{
					StatusCode: http.StatusNotFound,
					Body:       body,
				}, nil).Times(1)
			}},
			wantInfo: testdata.Ticket1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.GetByID(contextx.BackgroundWithLogger(s.logger), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByID() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *suiteHTTPClient) Test_httpclient_List() {
	type args struct {
		condition QueryTicketsCondition
		mock      func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo []*im.Ticket
		wantErr  bool
	}{
		{
			name: "list all tickets then error",
			args: args{condition: QueryTicketsCondition{
				Limit:  10,
				Offset: 0,
			}, mock: func() {
				req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/tasks?page=1&size=10", nil)
				s.client.EXPECT().Do(req).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "not found then nil",
			args: args{condition: QueryTicketsCondition{
				Limit:  10,
				Offset: 0,
			}, mock: func() {
				req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/tasks?page=1&size=10", nil)
				payload, _ := json.Marshal(response.Err.WithMessage("issue not found"))
				body := io.NopCloser(bytes.NewReader(payload))
				s.client.EXPECT().Do(req).Return(&http.Response{
					StatusCode: http.StatusNotFound,
					Body:       body,
				}, nil).Times(1)
			}},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "ok",
			args: args{condition: QueryTicketsCondition{
				Limit:  10,
				Offset: 0,
			}, mock: func() {
				req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/tasks?page=1&size=10", nil)
				payload, _ := json.Marshal(response.OK.WithData(struct {
					Total int          `json:"total"`
					List  []*im.Ticket `json:"list"`
				}{Total: 10, List: []*im.Ticket{testdata.Ticket1}}))
				body := io.NopCloser(bytes.NewReader(payload))
				s.client.EXPECT().Do(req).Return(&http.Response{
					StatusCode: http.StatusNotFound,
					Body:       body,
				}, nil).Times(1)
			}},
			wantInfo: []*im.Ticket{testdata.Ticket1},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.List(contextx.BackgroundWithLogger(s.logger), tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("List() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}
