// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "woman-center-be/internal/app/v1/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// ArticleFavoriteRepository is an autogenerated mock type for the ArticleFavoriteRepository type
type ArticleFavoriteRepository struct {
	mock.Mock
}

// AddFavoriteArticle provides a mock function with given fields: user, article
func (_m *ArticleFavoriteRepository) AddFavoriteArticle(user domain.Users, article *domain.Articles) error {
	ret := _m.Called(user, article)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Users, *domain.Articles) error); ok {
		r0 = rf(user, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFavoriteArticle provides a mock function with given fields: user, article
func (_m *ArticleFavoriteRepository) DeleteFavoriteArticle(user domain.Users, article *domain.Articles) error {
	ret := _m.Called(user, article)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Users, *domain.Articles) error); ok {
		r0 = rf(user, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewArticleFavoriteRepository creates a new instance of ArticleFavoriteRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewArticleFavoriteRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ArticleFavoriteRepository {
	mock := &ArticleFavoriteRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}