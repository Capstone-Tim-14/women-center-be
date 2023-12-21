// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "woman-center-be/internal/app/v1/models/domain"

	echo "github.com/labstack/echo/v4"

	exceptions "woman-center-be/utils/exceptions"

	mock "github.com/stretchr/testify/mock"

	requests "woman-center-be/internal/web/requests/v1"
)

// SpecialistService is an autogenerated mock type for the SpecialistService type
type SpecialistService struct {
	mock.Mock
}

// CreateSpecialist provides a mock function with given fields: ctx, request
func (_m *SpecialistService) CreateSpecialist(ctx echo.Context, request requests.SpecialistRequest) (*domain.Specialist, []exceptions.ValidationMessage, error) {
	ret := _m.Called(ctx, request)

	var r0 *domain.Specialist
	var r1 []exceptions.ValidationMessage
	var r2 error
	if rf, ok := ret.Get(0).(func(echo.Context, requests.SpecialistRequest) (*domain.Specialist, []exceptions.ValidationMessage, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, requests.SpecialistRequest) *domain.Specialist); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Specialist)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, requests.SpecialistRequest) []exceptions.ValidationMessage); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]exceptions.ValidationMessage)
		}
	}

	if rf, ok := ret.Get(2).(func(echo.Context, requests.SpecialistRequest) error); ok {
		r2 = rf(ctx, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteSpecialistById provides a mock function with given fields: ctx, id
func (_m *SpecialistService) DeleteSpecialistById(ctx echo.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetListSpecialist provides a mock function with given fields: ctx
func (_m *SpecialistService) GetListSpecialist(ctx echo.Context) ([]domain.Specialist, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Specialist
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context) ([]domain.Specialist, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(echo.Context) []domain.Specialist); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Specialist)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSpecialistService creates a new instance of SpecialistService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSpecialistService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SpecialistService {
	mock := &SpecialistService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}