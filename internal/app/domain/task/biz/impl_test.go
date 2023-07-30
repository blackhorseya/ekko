package biz

import (
	"reflect"
	"testing"

	taskB "github.com/blackhorseya/ekko/entity/domain/task/biz"
	taskM "github.com/blackhorseya/ekko/entity/domain/task/model"
	taskR "github.com/blackhorseya/ekko/internal/app/domain/task/biz/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteTester struct {
	suite.Suite

	logger *zap.Logger
	ctrl   *gomock.Controller
	repo   *taskR.MockIRepo
	biz    taskB.IBiz
}

func (s *SuiteTester) SetupTest() {
	s.logger = zap.NewExample()
	s.ctrl = gomock.NewController(s.T())
	s.repo = taskR.NewMockIRepo(s.ctrl)
	s.biz = NewImpl(s.repo)
}

func (s *SuiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	t.Parallel()
	t.Helper()

	suite.Run(t, new(SuiteTester))
}

func (s *SuiteTester) Test_impl_GetTicketByID() {
	ticket1 := &taskM.Ticket{
		Id:     "1",
		Title:  "title1",
		Status: taskM.TicketStatus_TICKET_STATUS_TODO,
	}

	type args struct {
		id   string
		mock func()
	}
	tests := []struct {
		name       string
		args       args
		wantTicket *taskM.Ticket
		wantErr    bool
	}{
		{
			name:       "empty id then error",
			args:       args{id: "   "},
			wantTicket: nil,
			wantErr:    true,
		},
		{
			name: "get by id then error",
			args: args{id: ticket1.Id, mock: func() {
				s.repo.EXPECT().GetTicketByID(gomock.Any(), ticket1.Id).Return(nil, errors.New("error")).Times(1)
			}},
			wantTicket: nil,
			wantErr:    true,
		},
		{
			name: "get by id if not found then error",
			args: args{id: ticket1.Id, mock: func() {
				s.repo.EXPECT().GetTicketByID(gomock.Any(), ticket1.Id).Return(nil, nil).Times(1)
			}},
			wantTicket: nil,
			wantErr:    true,
		},
		{
			name: "get by id then success",
			args: args{id: ticket1.Id, mock: func() {
				s.repo.EXPECT().GetTicketByID(gomock.Any(), ticket1.Id).Return(ticket1, nil).Times(1)
			}},
			wantTicket: ticket1,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTicket, err := s.biz.GetTicketByID(contextx.WithLogger(s.logger), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTicketByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTicket, tt.wantTicket) {
				t.Errorf("GetTicketByID() gotTicket = %v, want %v", gotTicket, tt.wantTicket)
			}
		})
	}
}

func (s *SuiteTester) Test_impl_ListTickets() {
	condB := taskB.ListTicketsCondition{Page: 1, Size: 10}
	condR := taskR.ListTicketsCondition{
		Limit:  10,
		Offset: 0,
	}
	ticket1 := &taskM.Ticket{
		Id:     "1",
		Title:  "title1",
		Status: taskM.TicketStatus_TICKET_STATUS_TODO,
	}
	tickets := []*taskM.Ticket{ticket1}

	type args struct {
		condition taskB.ListTicketsCondition
		mock      func()
	}
	tests := []struct {
		name        string
		args        args
		wantTickets []*taskM.Ticket
		wantTotal   int
		wantErr     bool
	}{
		{
			name:        "invalid page then error",
			args:        args{condition: taskB.ListTicketsCondition{Page: -1, Size: 10}},
			wantTickets: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name:        "invalid size then error",
			args:        args{condition: taskB.ListTicketsCondition{Page: 1, Size: -1}},
			wantTickets: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name: "list tickets then error",
			args: args{condition: condB, mock: func() {
				s.repo.EXPECT().ListTickets(gomock.Any(), condR).Return(nil, 0, errors.New("error")).Times(1)
			}},
			wantTickets: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name: "list tickets if not found then error",
			args: args{condition: condB, mock: func() {
				s.repo.EXPECT().ListTickets(gomock.Any(), condR).Return(nil, 10, nil).Times(1)
			}},
			wantTickets: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name: "list tickets then success",
			args: args{condition: condB, mock: func() {
				s.repo.EXPECT().ListTickets(gomock.Any(), condR).Return(tickets, 10, nil).Times(1)
			}},
			wantTickets: tickets,
			wantTotal:   10,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTickets, gotTotal, err := s.biz.ListTickets(contextx.WithLogger(s.logger), tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTickets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTickets, tt.wantTickets) {
				t.Errorf("ListTickets() gotTickets = %v, want %v", gotTickets, tt.wantTickets)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListTickets() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
