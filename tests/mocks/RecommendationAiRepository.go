// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "woman-center-be/internal/app/v1/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// RecommendationAiRepository is an autogenerated mock type for the RecommendationAiRepository type
type RecommendationAiRepository struct {
	mock.Mock
}

// FindAllHistoryRecommendationCareer provides a mock function with given fields: id
func (_m *RecommendationAiRepository) FindAllHistoryRecommendationCareer(id uint) (*domain.Users, error) {
	ret := _m.Called(id)

	var r0 *domain.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*domain.Users, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *domain.Users); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Users)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveHistoryRecommendationCareer provides a mock function with given fields: _a0
func (_m *RecommendationAiRepository) SaveHistoryRecommendationCareer(_a0 domain.HistoryRecommendationCareerAi) (*domain.HistoryRecommendationCareerAi, error) {
	ret := _m.Called(_a0)

	var r0 *domain.HistoryRecommendationCareerAi
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.HistoryRecommendationCareerAi) (*domain.HistoryRecommendationCareerAi, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(domain.HistoryRecommendationCareerAi) *domain.HistoryRecommendationCareerAi); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.HistoryRecommendationCareerAi)
		}
	}

	if rf, ok := ret.Get(1).(func(domain.HistoryRecommendationCareerAi) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRecommendationAiRepository creates a new instance of RecommendationAiRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRecommendationAiRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *RecommendationAiRepository {
	mock := &RecommendationAiRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
