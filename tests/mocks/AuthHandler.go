// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// AuthHandler is an autogenerated mock type for the AuthHandler type
type AuthHandler struct {
	mock.Mock
}

// AdminAuthHandler provides a mock function with given fields: _a0
func (_m *AuthHandler) AdminAuthHandler(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OauthCallbackGoogleHandler provides a mock function with given fields: _a0
func (_m *AuthHandler) OauthCallbackGoogleHandler(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OauthGoogleHandler provides a mock function with given fields: _a0
func (_m *AuthHandler) OauthGoogleHandler(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserAuthHandler provides a mock function with given fields: _a0
func (_m *AuthHandler) UserAuthHandler(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAuthHandler creates a new instance of AuthHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthHandler {
	mock := &AuthHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}