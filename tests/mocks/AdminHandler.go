// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// AdminHandler is an autogenerated mock type for the AdminHandler type
type AdminHandler struct {
	mock.Mock
}

// RegisterHandler provides a mock function with given fields: _a0
func (_m *AdminHandler) RegisterHandler(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAdminHandler creates a new instance of AdminHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdminHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdminHandler {
	mock := &AdminHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
