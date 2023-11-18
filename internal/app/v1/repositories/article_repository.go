package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(article *domain.Articles) (*domain.Articles, error)
	FindAllArticle(orderBy string) ([]domain.Articles, error)
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

func (repository *ArticleRepositoryImpl) FindAllArticle(orderBy string) ([]domain.Articles, error) {
	var articles []domain.Articles

	result := repository.db.Preload("Admin").Preload("Admin.Credential").Preload("Admin.Credential.Role").Preload("Counselors").Preload("Counselors.Credential").Preload("Counselors.Credential.Role")

	if orderBy != "" {
		result.Order("title " + orderBy).Find(&articles)
	} else {
		result.Find(&articles)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return articles, nil
}
