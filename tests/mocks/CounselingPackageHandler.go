// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// CounselingPackageHandler is an autogenerated mock type for the CounselingPackageHandler type
type CounselingPackageHandler struct {
	mock.Mock
}

// CreatePackage provides a mock function with given fields: ctx
func (_m *CounselingPackageHandler) CreatePackage(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePackage provides a mock function with given fields: ctx
func (_m *CounselingPackageHandler) DeletePackage(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByTitle provides a mock function with given fields: ctx
func (_m *CounselingPackageHandler) FindByTitle(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPackage provides a mock function with given fields: ctx
func (_m *CounselingPackageHandler) GetAllPackage(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePackage provides a mock function with given fields: ctx
func (_m *CounselingPackageHandler) UpdatePackage(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCounselingPackageHandler creates a new instance of CounselingPackageHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCounselingPackageHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *CounselingPackageHandler {
	mock := &CounselingPackageHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}