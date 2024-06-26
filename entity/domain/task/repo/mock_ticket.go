// Code generated by MockGen. DO NOT EDIT.
// Source: ticket.go

// Package repo is a generated GoMock package.
package repo

import (
	reflect "reflect"

	model "github.com/blackhorseya/ekko/entity/domain/task/model"
	contextx "github.com/blackhorseya/ekko/pkg/contextx"
	gomock "go.uber.org/mock/gomock"
)

// MockITicketRepo is a mock of ITicketRepo interface.
type MockITicketRepo struct {
	ctrl     *gomock.Controller
	recorder *MockITicketRepoMockRecorder
}

// MockITicketRepoMockRecorder is the mock recorder for MockITicketRepo.
type MockITicketRepoMockRecorder struct {
	mock *MockITicketRepo
}

// NewMockITicketRepo creates a new mock instance.
func NewMockITicketRepo(ctrl *gomock.Controller) *MockITicketRepo {
	mock := &MockITicketRepo{ctrl: ctrl}
	mock.recorder = &MockITicketRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITicketRepo) EXPECT() *MockITicketRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockITicketRepo) Create(ctx contextx.Contextx, ticket *model.Ticket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, ticket)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockITicketRepoMockRecorder) Create(ctx, ticket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockITicketRepo)(nil).Create), ctx, ticket)
}

// Delete mocks base method.
func (m *MockITicketRepo) Delete(ctx contextx.Contextx, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockITicketRepoMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITicketRepo)(nil).Delete), ctx, id)
}

// GetByID mocks base method.
func (m *MockITicketRepo) GetByID(ctx contextx.Contextx, id string) (*model.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*model.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockITicketRepoMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockITicketRepo)(nil).GetByID), ctx, id)
}

// List mocks base method.
func (m *MockITicketRepo) List(ctx contextx.Contextx, condition ListCondition) ([]*model.Ticket, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, condition)
	ret0, _ := ret[0].([]*model.Ticket)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockITicketRepoMockRecorder) List(ctx, condition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockITicketRepo)(nil).List), ctx, condition)
}

// Update mocks base method.
func (m *MockITicketRepo) Update(ctx contextx.Contextx, ticket *model.Ticket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, ticket)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockITicketRepoMockRecorder) Update(ctx, ticket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockITicketRepo)(nil).Update), ctx, ticket)
}
