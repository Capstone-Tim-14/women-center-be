package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type ArticlehasTagRepository interface {
	AddTag(article *domain.Articles, tag domain.Tag_Article) error
	RemoveTagById(article *domain.Articles, tag domain.Tag_Article) error
}

type ArticlehasTagRepositoryImpl struct {
	db *gorm.DB
}

func NewArticlehasTagRepository(db *gorm.DB) ArticlehasTagRepository {
	return &ArticlehasTagRepositoryImpl{
		db: db,
	}
}

func (repository *ArticlehasTagRepositoryImpl) AddTag(article *domain.Articles, tag domain.Tag_Article) error {
	result := repository.db.Model(&article).Association("Tag_Article").Append(tag)
	if result != nil {
		return result
	}

	return nil

}

func (repository *ArticlehasTagRepositoryImpl) RemoveTagById(article *domain.Articles, tag domain.Tag_Article) error {
	result := repository.db.Model(&article).Association("Tag").Delete(tag)
	if result != nil {
		return result
	}
	return nil
}
