// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "woman-center-be/internal/app/v1/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// JobTypeRepository is an autogenerated mock type for the JobTypeRepository type
type JobTypeRepository struct {
	mock.Mock
}

// CreateJobType provides a mock function with given fields: tag
func (_m *JobTypeRepository) CreateJobType(tag *domain.Job_Type) (*domain.Job_Type, error) {
	ret := _m.Called(tag)

	var r0 *domain.Job_Type
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Job_Type) (*domain.Job_Type, error)); ok {
		return rf(tag)
	}
	if rf, ok := ret.Get(0).(func(*domain.Job_Type) *domain.Job_Type); ok {
		r0 = rf(tag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Job_Type)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.Job_Type) error); ok {
		r1 = rf(tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobTypeByName provides a mock function with given fields: name
func (_m *JobTypeRepository) FindJobTypeByName(name string) (*domain.Job_Type, error) {
	ret := _m.Called(name)

	var r0 *domain.Job_Type
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Job_Type, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Job_Type); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Job_Type)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowAllJobType provides a mock function with given fields:
func (_m *JobTypeRepository) ShowAllJobType() ([]domain.Job_Type, error) {
	ret := _m.Called()

	var r0 []domain.Job_Type
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Job_Type, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Job_Type); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Job_Type)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJobTypeRepository creates a new instance of JobTypeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJobTypeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *JobTypeRepository {
	mock := &JobTypeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
