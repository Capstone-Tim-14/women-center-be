package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type ArticlehasTagRepository interface {
	AddTag(article domain.Articles, tag *domain.Tag_Article) error
	RemoveTag(article domain.Articles, tag []domain.Tag_Article) error
}

type ArticlehasTagRepositoryImpl struct {
	Db *gorm.DB
}

func NewArticlehasTagRepository(db *gorm.DB) ArticlehasTagRepository {
	return &ArticlehasTagRepositoryImpl{
		Db: db,
	}
}

func (repository *ArticlehasTagRepositoryImpl) AddTag(article domain.Articles, tag *domain.Tag_Article) error {
	result := repository.Db.Model(&article).Association("Tags").Append(tag)
	if result != nil {
		return result
	}

	return nil

}

func (repository *ArticlehasTagRepositoryImpl) RemoveTag(article domain.Articles, tag []domain.Tag_Article) error {

	Transaction := repository.Db.Begin()

	result := Transaction.Model(&article).Association("Tags").Delete(tag)
	if result != nil {
		Transaction.Rollback()
		return result
	}

	Transaction.Commit()

	return nil
}
