// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	exceptions "woman-center-be/utils/exceptions"

	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"

	requests "woman-center-be/internal/web/requests/v1"
)

// ScheduleService is an autogenerated mock type for the ScheduleService type
type ScheduleService struct {
	mock.Mock
}

// CreateSchedule provides a mock function with given fields: ctx, request
func (_m *ScheduleService) CreateSchedule(ctx echo.Context, request []requests.CounselingScheduleRequest) ([]exceptions.ValidationMessage, error) {
	ret := _m.Called(ctx, request)

	var r0 []exceptions.ValidationMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, []requests.CounselingScheduleRequest) ([]exceptions.ValidationMessage, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, []requests.CounselingScheduleRequest) []exceptions.ValidationMessage); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]exceptions.ValidationMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, []requests.CounselingScheduleRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteScheduletById provides a mock function with given fields: ctx, id
func (_m *ScheduleService) DeleteScheduletById(ctx echo.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateScheduleById provides a mock function with given fields: ctx, request
func (_m *ScheduleService) UpdateScheduleById(ctx echo.Context, request requests.CounselingScheduleRequest) ([]exceptions.ValidationMessage, error) {
	ret := _m.Called(ctx, request)

	var r0 []exceptions.ValidationMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, requests.CounselingScheduleRequest) ([]exceptions.ValidationMessage, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, requests.CounselingScheduleRequest) []exceptions.ValidationMessage); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]exceptions.ValidationMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, requests.CounselingScheduleRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewScheduleService creates a new instance of ScheduleService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewScheduleService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ScheduleService {
	mock := &ScheduleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}