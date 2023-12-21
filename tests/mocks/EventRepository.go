// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "woman-center-be/internal/app/v1/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// EventRepository is an autogenerated mock type for the EventRepository type
type EventRepository struct {
	mock.Mock
}

// CreateEvent provides a mock function with given fields: event
func (_m *EventRepository) CreateEvent(event *domain.Event) (*domain.Event, error) {
	ret := _m.Called(event)

	var r0 *domain.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Event) (*domain.Event, error)); ok {
		return rf(event)
	}
	if rf, ok := ret.Get(0).(func(*domain.Event) *domain.Event); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.Event) error); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllEvent provides a mock function with given fields:
func (_m *EventRepository) FindAllEvent() ([]domain.Event, error) {
	ret := _m.Called()

	var r0 []domain.Event
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Event, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Event)
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
func (_m *EventRepository) FindById(id int) (*domain.Event, error) {
	ret := _m.Called(id)

	var r0 *domain.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*domain.Event, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *domain.Event); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindDetailEvent provides a mock function with given fields: id
func (_m *EventRepository) FindDetailEvent(id int) (*domain.Event, error) {
	ret := _m.Called(id)

	var r0 *domain.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*domain.Event, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *domain.Event); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEvent provides a mock function with given fields: id, event
func (_m *EventRepository) UpdateEvent(id int, event *domain.Event) error {
	ret := _m.Called(id, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *domain.Event) error); ok {
		r0 = rf(id, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEventRepository creates a new instance of EventRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventRepository {
	mock := &EventRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
