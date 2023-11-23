package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type TagRepository interface {
	CreateTag(tag *domain.Tag_Article) (*domain.Tag_Article, error)
	FindById(id int) (*domain.Tag_Article, error)
	FindTagByName(name string) (*domain.Tag_Article, error)
	FindAllTags() ([]domain.Tag_Article, error)
	DeleteTagById(tagID int) error
}

type TagRepositoryImpl struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &TagRepositoryImpl{
		db: db,
	}
}

func (repository *TagRepositoryImpl) CreateTag(tag *domain.Tag_Article) (*domain.Tag_Article, error) {
	result := repository.db.Create(&tag)
	if result.Error != nil {
		return nil, result.Error
	}

	return tag, nil
}

func (repository *TagRepositoryImpl) FindById(id int) (*domain.Tag_Article, error) {
	tag := domain.Tag_Article{}

	result := repository.db.Where("id = ?", id).First(&tag)
	if result.Error != nil {
		return nil, result.Error
	}

	return &tag, nil
}

func (repository *TagRepositoryImpl) FindTagByName(name string) (*domain.Tag_Article, error) {
	tag := domain.Tag_Article{}

	result := repository.db.Where("name = ?", name).First(&tag)
	if result.Error != nil {
		return nil, result.Error
	}

	return &tag, nil
}

func (repository *TagRepositoryImpl) FindAllTags() ([]domain.Tag_Article, error) {
	var tags []domain.Tag_Article

	result := repository.db.Find(&tags)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Tags is empty")
	}

	return tags, nil
}

func (repository *TagRepositoryImpl) DeleteTagById(tagID int) error {
	result := repository.db.Delete(&domain.Tag_Article{}, tagID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
