package biz

import (
	"reflect"
	"testing"

	issueB "github.com/blackhorseya/ekko/entity/domain/issue/biz"
	issueM "github.com/blackhorseya/ekko/entity/domain/issue/model"
	"github.com/blackhorseya/ekko/internal/app/domain/issue/biz/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/genx"
	"github.com/blackhorseya/ekko/test/testdata"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite
	logger    *zap.Logger
	ctrl      *gomock.Controller
	generator *genx.MockGenerator
	repo      *repo.MockIRepo
	biz       issueB.IBiz
}

func (s *suiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.ctrl = gomock.NewController(s.T())
	s.generator = genx.NewMockGenerator(s.ctrl)
	s.repo = repo.NewMockIRepo(s.ctrl)
	s.biz = CreateBiz(s.repo, s.generator)
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_GetByID() {
	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *issueM.Ticket
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{id: testdata.Ticket1.Id, mock: func() {
				s.repo.EXPECT().GetByID(gomock.Any(), testdata.Ticket1.Id).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then ok",
			args: args{id: testdata.Ticket1.Id, mock: func() {
				s.repo.EXPECT().GetByID(gomock.Any(), testdata.Ticket1.Id).Return(testdata.Ticket1, nil).Times(1)
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

			gotInfo, err := s.biz.GetByID(contextx.WithLogger(s.logger), tt.args.id)
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

func (s *suiteTester) Test_impl_List() {
	type args struct {
		page int
		size int
		mock func()
	}
	tests := []struct {
		name      string
		args      args
		wantInfo  []*issueM.Ticket
		wantTotal int
		wantErr   bool
	}{
		{
			name:      "invalid page then error",
			args:      args{page: -1, size: 10},
			wantInfo:  nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name:      "invalid size then error",
			args:      args{page: 1, size: -10},
			wantInfo:  nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "list tickets then error",
			args: args{page: 2, size: 10, mock: func() {
				s.repo.EXPECT().List(gomock.Any(), repo.QueryTicketsCondition{
					Limit:  10,
					Offset: 10,
				}).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo:  nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "if not exists then ok",
			args: args{page: 2, size: 10, mock: func() {
				s.repo.EXPECT().List(gomock.Any(), repo.QueryTicketsCondition{
					Limit:  10,
					Offset: 10,
				}).Return(nil, nil).Times(1)
			}},
			wantInfo:  nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "count all tickets then error",
			args: args{page: 2, size: 10, mock: func() {
				condition := repo.QueryTicketsCondition{
					Limit:  10,
					Offset: 10,
				}
				s.repo.EXPECT().List(gomock.Any(), condition).Return([]*issueM.Ticket{testdata.Ticket1}, nil).Times(1)

				s.repo.EXPECT().Count(gomock.Any(), condition).Return(0, errors.New("error")).Times(1)
			}},
			wantInfo:  nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "ok",
			args: args{page: 2, size: 10, mock: func() {
				condition := repo.QueryTicketsCondition{
					Limit:  10,
					Offset: 10,
				}
				s.repo.EXPECT().List(gomock.Any(), condition).Return([]*issueM.Ticket{testdata.Ticket1}, nil).Times(1)

				s.repo.EXPECT().Count(gomock.Any(), condition).Return(10, nil).Times(1)
			}},
			wantInfo:  []*issueM.Ticket{testdata.Ticket1},
			wantTotal: 10,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, gotTotal, err := s.biz.List(contextx.WithLogger(s.logger), tt.args.page, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("List() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("List() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func (s *suiteTester) Test_impl_Create() {
	type args struct {
		title string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *issueM.Ticket
		wantErr  bool
	}{
		{
			name:     "invalid title then error",
			args:     args{title: "   "},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "create a ticket then error",
			args: args{title: testdata.Ticket1.Title, mock: func() {
				s.generator.EXPECT().Int64().Return(testdata.Ticket1.Id).Times(1)
				created := &issueM.Ticket{
					Id:        testdata.Ticket1.Id,
					Title:     testdata.Ticket1.Title,
					Status:    issueM.TicketStatus_TICKET_STATUS_TODO,
					CreatedAt: nil,
					UpdatedAt: nil,
				}
				s.repo.EXPECT().Create(gomock.Any(), created).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "ok",
			args: args{title: testdata.Ticket1.Title, mock: func() {
				s.generator.EXPECT().Int64().Return(testdata.Ticket1.Id).Times(1)
				created := &issueM.Ticket{
					Id:        testdata.Ticket1.Id,
					Title:     testdata.Ticket1.Title,
					Status:    issueM.TicketStatus_TICKET_STATUS_TODO,
					CreatedAt: nil,
					UpdatedAt: nil,
				}
				s.repo.EXPECT().Create(gomock.Any(), created).Return(testdata.Ticket1, nil).Times(1)
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

			gotInfo, err := s.biz.Create(contextx.WithLogger(s.logger), tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Create() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *suiteTester) Test_impl_UpdateStatus() {
	type args struct {
		id     int64
		status issueM.TicketStatus
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *issueM.Ticket
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{id: testdata.Ticket1.Id, status: issueM.TicketStatus_TICKET_STATUS_DONE, mock: func() {
				s.repo.EXPECT().GetByID(gomock.Any(), testdata.Ticket1.Id).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "if not exists then error",
			args: args{id: testdata.Ticket1.Id, status: issueM.TicketStatus_TICKET_STATUS_DONE, mock: func() {
				s.repo.EXPECT().GetByID(gomock.Any(), testdata.Ticket1.Id).Return(nil, nil).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "update status then error",
			args: args{id: testdata.Ticket1.Id, status: issueM.TicketStatus_TICKET_STATUS_DONE, mock: func() {
				s.repo.EXPECT().GetByID(gomock.Any(), testdata.Ticket1.Id).Return(testdata.Ticket1, nil).Times(1)

				s.repo.EXPECT().Update(gomock.Any(), testdata.Ticket1).Return(errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "ok",
			args: args{id: testdata.Ticket1.Id, status: issueM.TicketStatus_TICKET_STATUS_DONE, mock: func() {
				s.repo.EXPECT().GetByID(gomock.Any(), testdata.Ticket1.Id).Return(testdata.Ticket1, nil).Times(1)

				s.repo.EXPECT().Update(gomock.Any(), testdata.Ticket1).Return(nil).Times(1)
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

			gotInfo, err := s.biz.UpdateStatus(contextx.WithLogger(s.logger), tt.args.id, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("UpdateStatus() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *suiteTester) Test_impl_Delete() {
	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete by id then error",
			args: args{id: testdata.Ticket1.Id, mock: func() {
				s.repo.EXPECT().DeleteByID(gomock.Any(), testdata.Ticket1.Id).Return(errors.New("error")).Times(1)
			}},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{id: testdata.Ticket1.Id, mock: func() {
				s.repo.EXPECT().DeleteByID(gomock.Any(), testdata.Ticket1.Id).Return(nil).Times(1)
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.Delete(contextx.WithLogger(s.logger), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
