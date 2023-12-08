package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type ArticleFavoriteRepository interface {
	AddFavoriteArticle(user domain.Users, article *domain.Articles) error
	//DeleteFavoriteArticle(user domain.Users, article *domain.Articles) error
}

type ArticleFavoriteRepositoryImpl struct {
	Db *gorm.DB
}

func NewArticleFavoriteRepository(db *gorm.DB) ArticleFavoriteRepository {
	return &ArticleFavoriteRepositoryImpl{
		Db: db,
	}
}

func (repo *ArticleFavoriteRepositoryImpl) AddFavoriteArticle(user domain.Users, article *domain.Articles) error {
	result := repo.Db.Model(&user).Association("Articles").Append(article)
	if result != nil {
		return result
	}

	return nil
}

func (repo *ArticleFavoriteRepositoryImpl) DeleteFavoriteArticle(user domain.Users, article *domain.Articles) error {
	result := repo.Db.Model(&user).Association("Articles").Delete(article)
	if result != nil {
		return result
	}

	return nil
}
