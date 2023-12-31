// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// CounselingSessionHandlers is an autogenerated mock type for the CounselingSessionHandlers type
type CounselingSessionHandlers struct {
	mock.Mock
}

// CounselingSessionDetailHandler provides a mock function with given fields: _a0
func (_m *CounselingSessionHandlers) CounselingSessionDetailHandler(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListCounselingSessionHandler provides a mock function with given fields: _a0
func (_m *CounselingSessionHandlers) ListCounselingSessionHandler(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCounselingSessionHandlers creates a new instance of CounselingSessionHandlers. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCounselingSessionHandlers(t interface {
	mock.TestingT
	Cleanup(func())
}) *CounselingSessionHandlers {
	mock := &CounselingSessionHandlers{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
