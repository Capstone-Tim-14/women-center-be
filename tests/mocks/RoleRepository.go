// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "woman-center-be/internal/app/v1/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// RoleRepository is an autogenerated mock type for the RoleRepository type
type RoleRepository struct {
	mock.Mock
}

// CreateRole provides a mock function with given fields: role
func (_m *RoleRepository) CreateRole(role *domain.Roles) (*domain.Roles, error) {
	ret := _m.Called(role)

	var r0 *domain.Roles
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Roles) (*domain.Roles, error)); ok {
		return rf(role)
	}
	if rf, ok := ret.Get(0).(func(*domain.Roles) *domain.Roles); ok {
		r0 = rf(role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Roles)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.Roles) error); ok {
		r1 = rf(role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteById provides a mock function with given fields: id
func (_m *RoleRepository) DeleteById(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *RoleRepository) FindAll() ([]domain.Roles, error) {
	ret := _m.Called()

	var r0 []domain.Roles
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Roles, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Roles); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Roles)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *RoleRepository) FindById(id int) (*domain.Roles, error) {
	ret := _m.Called(id)

	var r0 *domain.Roles
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*domain.Roles, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *domain.Roles); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Roles)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: name
func (_m *RoleRepository) FindByName(name string) (*domain.Roles, error) {
	ret := _m.Called(name)

	var r0 *domain.Roles
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Roles, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Roles); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Roles)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRoleRepository creates a new instance of RoleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRoleRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *RoleRepository {
	mock := &RoleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
