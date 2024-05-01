// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GerarTokenInterface is an autogenerated mock type for the GerarTokenInterface type
type GerarTokenInterface struct {
	mock.Mock
}

// GenerateTokenJWT provides a mock function with given fields:
func (_m *GerarTokenInterface) GenerateTokenJWT() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GenerateTokenJWT")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGerarTokenInterface creates a new instance of GerarTokenInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGerarTokenInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *GerarTokenInterface {
	mock := &GerarTokenInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
