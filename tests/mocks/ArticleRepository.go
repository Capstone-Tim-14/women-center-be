// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "woman-center-be/internal/app/v1/models/domain"

	mock "github.com/stretchr/testify/mock"

	query "woman-center-be/pkg/query"
)

// ArticleRepository is an autogenerated mock type for the ArticleRepository type
type ArticleRepository struct {
	mock.Mock
}

// CreateArticle provides a mock function with given fields: article
func (_m *ArticleRepository) CreateArticle(article *domain.Articles) (*domain.Articles, error) {
	ret := _m.Called(article)

	var r0 *domain.Articles
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Articles) (*domain.Articles, error)); ok {
		return rf(article)
	}
	if rf, ok := ret.Get(0).(func(*domain.Articles) *domain.Articles); ok {
		r0 = rf(article)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Articles)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.Articles) error); ok {
		r1 = rf(article)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteArticleById provides a mock function with given fields: id
func (_m *ArticleRepository) DeleteArticleById(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindActiveArticleBySlug provides a mock function with given fields: slug
func (_m *ArticleRepository) FindActiveArticleBySlug(slug string) (*domain.Articles, error) {
	ret := _m.Called(slug)

	var r0 *domain.Articles
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Articles, error)); ok {
		return rf(slug)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Articles); ok {
		r0 = rf(slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Articles)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllArticle provides a mock function with given fields: _a0, _a1, _a2
func (_m *ArticleRepository) FindAllArticle(_a0 string, _a1 string, _a2 query.Pagination) ([]domain.Articles, *query.Pagination, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []domain.Articles
	var r1 *query.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string, query.Pagination) ([]domain.Articles, *query.Pagination, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(string, string, query.Pagination) []domain.Articles); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Articles)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, query.Pagination) *query.Pagination); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*query.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string, query.Pagination) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindArticleCounselor provides a mock function with given fields: id
func (_m *ArticleRepository) FindArticleCounselor(id int) ([]domain.Articles, *domain.ArticleStatusCount, error) {
	ret := _m.Called(id)

	var r0 []domain.Articles
	var r1 *domain.ArticleStatusCount
	var r2 error
	if rf, ok := ret.Get(0).(func(int) ([]domain.Articles, *domain.ArticleStatusCount, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) []domain.Articles); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Articles)
		}
	}

	if rf, ok := ret.Get(1).(func(int) *domain.ArticleStatusCount); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.ArticleStatusCount)
		}
	}

	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindArticleCounselorById provides a mock function with given fields: counselorID, articleID
func (_m *ArticleRepository) FindArticleCounselorById(counselorID int, articleID int) (*domain.Articles, error) {
	ret := _m.Called(counselorID, articleID)

	var r0 *domain.Articles
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (*domain.Articles, error)); ok {
		return rf(counselorID, articleID)
	}
	if rf, ok := ret.Get(0).(func(int, int) *domain.Articles); ok {
		r0 = rf(counselorID, articleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Articles)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(counselorID, articleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *ArticleRepository) FindById(id int) (*domain.Articles, error) {
	ret := _m.Called(id)

	var r0 *domain.Articles
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*domain.Articles, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *domain.Articles); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Articles)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBySlug provides a mock function with given fields: slug
func (_m *ArticleRepository) FindBySlug(slug string) (*domain.Articles, error) {
	ret := _m.Called(slug)

	var r0 *domain.Articles
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Articles, error)); ok {
		return rf(slug)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Articles); ok {
		r0 = rf(slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Articles)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByTitle provides a mock function with given fields: title
func (_m *ArticleRepository) FindByTitle(title string) (*domain.Articles, error) {
	ret := _m.Called(title)

	var r0 *domain.Articles
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Articles, error)); ok {
		return rf(title)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Articles); ok {
		r0 = rf(title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Articles)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSlugForFavorite provides a mock function with given fields: slug
func (_m *ArticleRepository) FindSlugForFavorite(slug string) (*domain.Articles, error) {
	ret := _m.Called(slug)

	var r0 *domain.Articles
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Articles, error)); ok {
		return rf(slug)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Articles); ok {
		r0 = rf(slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Articles)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestArticleForUser provides a mock function with given fields:
func (_m *ArticleRepository) GetLatestArticleForUser() (*domain.Articles, error) {
	ret := _m.Called()

	var r0 *domain.Articles
	var r1 error
	if rf, ok := ret.Get(0).(func() (*domain.Articles, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *domain.Articles); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Articles)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListArticleForUser provides a mock function with given fields:
func (_m *ArticleRepository) GetListArticleForUser() ([]domain.Articles, error) {
	ret := _m.Called()

	var r0 []domain.Articles
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Articles, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Articles); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Articles)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateArticle provides a mock function with given fields: id, article
func (_m *ArticleRepository) UpdateArticle(id int, article *domain.Articles) error {
	ret := _m.Called(id, article)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *domain.Articles) error); ok {
		r0 = rf(id, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatusArticle provides a mock function with given fields: slug, status
func (_m *ArticleRepository) UpdateStatusArticle(slug string, status string) error {
	ret := _m.Called(slug, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(slug, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewArticleRepository creates a new instance of ArticleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewArticleRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ArticleRepository {
	mock := &ArticleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
