// Code generated by mockery v2.14.1. DO NOT EDIT.

package biz

import (
	contextx "github.com/blackhorseya/todo-app/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	model "github.com/blackhorseya/todo-app/pkg/entity/domain/todo/model"
)

// MockIBiz is an autogenerated mock type for the IBiz type
type MockIBiz struct {
	mock.Mock
}

// ChangeTitle provides a mock function with given fields: ctx, id, title
func (_m *MockIBiz) ChangeTitle(ctx contextx.Contextx, id int64, title string) (*model.Task, error) {
	ret := _m.Called(ctx, id, title)

	var r0 *model.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64, string) *model.Task); ok {
		r0 = rf(ctx, id, title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int64, string) error); ok {
		r1 = rf(ctx, id, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, title
func (_m *MockIBiz) Create(ctx contextx.Contextx, title string) (*model.Task, error) {
	ret := _m.Called(ctx, title)

	var r0 *model.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *model.Task); ok {
		r0 = rf(ctx, title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockIBiz) Delete(ctx contextx.Contextx, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *MockIBiz) GetByID(ctx contextx.Contextx, id int64) (*model.Task, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64) *model.Task); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, page, size
func (_m *MockIBiz) List(ctx contextx.Contextx, page int, size int) ([]*model.Task, int, error) {
	ret := _m.Called(ctx, page, size)

	var r0 []*model.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int, int) []*model.Task); ok {
		r0 = rf(ctx, page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Task)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int, int) int); ok {
		r1 = rf(ctx, page, size)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(contextx.Contextx, int, int) error); ok {
		r2 = rf(ctx, page, size)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Liveness provides a mock function with given fields: ctx
func (_m *MockIBiz) Liveness(ctx contextx.Contextx) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Readiness provides a mock function with given fields: ctx
func (_m *MockIBiz) Readiness(ctx contextx.Contextx) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatus provides a mock function with given fields: ctx, id, status
func (_m *MockIBiz) UpdateStatus(ctx contextx.Contextx, id int64, status model.TaskStatus) (*model.Task, error) {
	ret := _m.Called(ctx, id, status)

	var r0 *model.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64, model.TaskStatus) *model.Task); ok {
		r0 = rf(ctx, id, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int64, model.TaskStatus) error); ok {
		r1 = rf(ctx, id, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockIBiz interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIBiz creates a new instance of MockIBiz. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIBiz(t mockConstructorTestingTNewMockIBiz) *MockIBiz {
	mock := &MockIBiz{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}