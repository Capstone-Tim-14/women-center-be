// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "woman-center-be/internal/app/v1/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// ArticlehasTagRepository is an autogenerated mock type for the ArticlehasTagRepository type
type ArticlehasTagRepository struct {
	mock.Mock
}

// AddTag provides a mock function with given fields: article, tag
func (_m *ArticlehasTagRepository) AddTag(article domain.Articles, tag *domain.Tag_Article) error {
	ret := _m.Called(article, tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Articles, *domain.Tag_Article) error); ok {
		r0 = rf(article, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveTag provides a mock function with given fields: article, tag
func (_m *ArticlehasTagRepository) RemoveTag(article domain.Articles, tag []domain.Tag_Article) error {
	ret := _m.Called(article, tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Articles, []domain.Tag_Article) error); ok {
		r0 = rf(article, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewArticlehasTagRepository creates a new instance of ArticlehasTagRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewArticlehasTagRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ArticlehasTagRepository {
	mock := &ArticlehasTagRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}