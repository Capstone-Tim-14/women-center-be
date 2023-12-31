// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// TagHandler is an autogenerated mock type for the TagHandler type
type TagHandler struct {
	mock.Mock
}

// CreateTag provides a mock function with given fields: ctx
func (_m *TagHandler) CreateTag(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTagByIdHandler provides a mock function with given fields: _a0
func (_m *TagHandler) DeleteTagByIdHandler(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTagsHandler provides a mock function with given fields: ctx
func (_m *TagHandler) GetAllTagsHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTagHandler creates a new instance of TagHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTagHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *TagHandler {
	mock := &TagHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
