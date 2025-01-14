// Code generated by mockery v2.42.2. DO NOT EDIT.

package tests

import (
	context "context"

	api "github.com/walletera/payments-types/api"

	mock "github.com/stretchr/testify/mock"
)

// MockSecuritySource is an autogenerated mock type for the SecuritySource type
type MockSecuritySource struct {
	mock.Mock
}

type MockSecuritySource_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSecuritySource) EXPECT() *MockSecuritySource_Expecter {
	return &MockSecuritySource_Expecter{mock: &_m.Mock}
}

// BearerAuth provides a mock function with given fields: ctx, operationName
func (_m *MockSecuritySource) BearerAuth(ctx context.Context, operationName string) (api.BearerAuth, error) {
	ret := _m.Called(ctx, operationName)

	if len(ret) == 0 {
		panic("no return value specified for BearerAuth")
	}

	var r0 api.BearerAuth
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (api.BearerAuth, error)); ok {
		return rf(ctx, operationName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) api.BearerAuth); ok {
		r0 = rf(ctx, operationName)
	} else {
		r0 = ret.Get(0).(api.BearerAuth)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, operationName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSecuritySource_BearerAuth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BearerAuth'
type MockSecuritySource_BearerAuth_Call struct {
	*mock.Call
}

// BearerAuth is a helper method to define mock.On call
//   - ctx context.Context
//   - operationName string
func (_e *MockSecuritySource_Expecter) BearerAuth(ctx interface{}, operationName interface{}) *MockSecuritySource_BearerAuth_Call {
	return &MockSecuritySource_BearerAuth_Call{Call: _e.mock.On("BearerAuth", ctx, operationName)}
}

func (_c *MockSecuritySource_BearerAuth_Call) Run(run func(ctx context.Context, operationName string)) *MockSecuritySource_BearerAuth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSecuritySource_BearerAuth_Call) Return(_a0 api.BearerAuth, _a1 error) *MockSecuritySource_BearerAuth_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSecuritySource_BearerAuth_Call) RunAndReturn(run func(context.Context, string) (api.BearerAuth, error)) *MockSecuritySource_BearerAuth_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSecuritySource creates a new instance of MockSecuritySource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSecuritySource(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSecuritySource {
	mock := &MockSecuritySource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
