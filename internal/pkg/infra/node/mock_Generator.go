// Code generated by mockery v2.14.0. DO NOT EDIT.

package node

import (
	snowflake "github.com/bwmarrin/snowflake"
	mock "github.com/stretchr/testify/mock"
)

// MockGenerator is an autogenerated mock type for the Generator type
type MockGenerator struct {
	mock.Mock
}

// Generate provides a mock function with given fields:
func (_m *MockGenerator) Generate() snowflake.ID {
	ret := _m.Called()

	var r0 snowflake.ID
	if rf, ok := ret.Get(0).(func() snowflake.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(snowflake.ID)
	}

	return r0
}

type mockConstructorTestingTNewMockGenerator interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockGenerator creates a new instance of MockGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockGenerator(t mockConstructorTestingTNewMockGenerator) *MockGenerator {
	mock := &MockGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}