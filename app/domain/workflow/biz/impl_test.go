package biz

import (
	"reflect"
	"testing"
	"time"

	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/agg"
	"github.com/blackhorseya/ekko/entity/domain/workflow/biz"
	"github.com/blackhorseya/ekko/entity/domain/workflow/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type suiteTester struct {
	suite.Suite

	ctrl   *gomock.Controller
	issues *repo.MockIIssueRepo
	biz    biz.IWorkflowBiz
}

func (s *suiteTester) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.issues = repo.NewMockIIssueRepo(s.ctrl)
	s.biz = NewWorkflowBiz(s.issues)
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_CreateTodo() {
	user1 := &idM.User{
		ID: "user1",
	}
	issue1 := &agg.Issue{
		Ticket: &model.Ticket{
			ID:        "",
			Title:     "issue1",
			Completed: false,
			OwnerID:   user1.ID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}

	type args struct {
		ctx   contextx.Contextx
		who   *idM.User
		title string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantItem *agg.Issue
		wantErr  bool
	}{
		{
			name: "create todo then ok",
			args: args{who: user1, title: issue1.Title, mock: func() {
				s.issues.EXPECT().Create(gomock.Any(), issue1).Return(nil).Times(1)
			}},
			wantItem: issue1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotItem, err := s.biz.CreateTodo(tt.args.ctx, tt.args.who, tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("CreateTodo() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func (s *suiteTester) Test_impl_ListTodos() {
	user1 := &idM.User{
		ID: "user1",
	}
	issue1 := &agg.Issue{
		Ticket: &model.Ticket{
			ID:        "",
			Title:     "issue1",
			Completed: false,
			OwnerID:   user1.ID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}

	type args struct {
		ctx  contextx.Contextx
		who  *idM.User
		opts biz.ListTodosOptions
		mock func()
	}
	tests := []struct {
		name      string
		args      args
		wantItems []*agg.Issue
		wantTotal int
		wantErr   bool
	}{
		{
			name: "list todos then ok",
			args: args{who: user1, opts: biz.ListTodosOptions{Page: 1, Size: 5}, mock: func() {
				s.issues.EXPECT().List(gomock.Any(), repo.ListIssueOptions{
					OwnerID: user1.ID,
					Limit:   5,
					Offset:  0,
				}).Return([]*agg.Issue{issue1}, 10, nil).Times(1)
			}},
			wantItems: []*agg.Issue{issue1},
			wantTotal: 10,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotItems, gotTotal, err := s.biz.ListTodos(tt.args.ctx, tt.args.who, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItems, tt.wantItems) {
				t.Errorf("ListTodos() gotItems = %v, want %v", gotItems, tt.wantItems)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListTodos() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
