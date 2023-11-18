package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(article *domain.Articles) (*domain.Articles, error)
	FindAllArticle() ([]domain.Articles, error)
	FindBySlug(slug string) (*domain.Articles, error)
	FindByTitle(title string) (*domain.Articles, error)
}

type ArticleRepositoryImpl struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &ArticleRepositoryImpl{
		db: db,
	}
}

func (repository *ArticleRepositoryImpl) CreateArticle(article *domain.Articles) (*domain.Articles, error) {
	result := repository.db.Create(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return article, nil

}

func (repository *ArticleRepositoryImpl) FindAllArticle() ([]domain.Articles, error) {
	articles := []domain.Articles{}
	result := repository.db.Preload("Admin").Preload("Admin.Credential").Preload("Admin.Credential.Role").Preload("Counselors").Preload("Counselors.Credential").Preload("Counselors.Credential.Role").Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}

	return articles, nil
}

func (repository *ArticleRepositoryImpl) FindBySlug(slug string) (*domain.Articles, error) {
	article := domain.Articles{}
	result := repository.db.Preload("Admin").Preload("Admin.Credential").Preload("Admin.Credential.Role").Preload("Counselors").Preload("Counselors.Credential").Preload("Counselors.Credential.Role").Where("slug = ?", slug).First(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return &article, nil
}

func (repository *ArticleRepositoryImpl) FindByTitle(title string) (*domain.Articles, error) {
	article := domain.Articles{}
	result := repository.db.Where("title = ?", title).First(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return &article, nil
}
