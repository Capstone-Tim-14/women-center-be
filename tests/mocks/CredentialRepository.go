// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "woman-center-be/internal/app/v1/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// CredentialRepository is an autogenerated mock type for the CredentialRepository type
type CredentialRepository struct {
	mock.Mock
}

// CheckEmailCredential provides a mock function with given fields: email
func (_m *CredentialRepository) CheckEmailCredential(email string) (*domain.Credentials, error) {
	ret := _m.Called(email)

	var r0 *domain.Credentials
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Credentials, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Credentials); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Credentials)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthUser provides a mock function with given fields: id, role
func (_m *CredentialRepository) GetAuthUser(id uint, role string) (*domain.Users, *domain.Counselors, error) {
	ret := _m.Called(id, role)

	var r0 *domain.Users
	var r1 *domain.Counselors
	var r2 error
	if rf, ok := ret.Get(0).(func(uint, string) (*domain.Users, *domain.Counselors, error)); ok {
		return rf(id, role)
	}
	if rf, ok := ret.Get(0).(func(uint, string) *domain.Users); ok {
		r0 = rf(id, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Users)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, string) *domain.Counselors); ok {
		r1 = rf(id, role)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.Counselors)
		}
	}

	if rf, ok := ret.Get(2).(func(uint, string) error); ok {
		r2 = rf(id, role)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewCredentialRepository creates a new instance of CredentialRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCredentialRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CredentialRepository {
	mock := &CredentialRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
