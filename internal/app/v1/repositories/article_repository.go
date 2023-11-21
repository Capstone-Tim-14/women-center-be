package repositories

import (
	"fmt"
	"strconv"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/pkg/query"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(article *domain.Articles) (*domain.Articles, error)
	FindAllArticle(string, query.Pagination) ([]domain.Articles, *query.Pagination, error)
	DeleteArticleById(id int) error
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

func (repository *ArticleRepositoryImpl) FindAllArticle(orderBy string, paginate query.Pagination) ([]domain.Articles, *query.Pagination, error) {
	var articles []domain.Articles

	result := repository.db.Scopes(query.Paginate(articles, &paginate, repository.db)).Preload("Admin").Preload("Admin.Credential").Preload("Admin.Credential.Role").Preload("Counselors").Preload("Counselors.Credential").Preload("Counselors.Credential.Role")

	if orderBy != "" {
		result.Order("title " + orderBy).Find(&articles)
	} else {
		result.Find(&articles)
	}

	if result.Error != nil {
		return nil, nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil, fmt.Errorf("Article is empty")
	}

	if paginate.Page <= 1 {
		paginate.PreviousPage = ""
	} else {
		paginate.PreviousPage = viper.GetString("MAIN_URL") + "/api/" + viper.GetString("API_VERSION") + "/admin/articles?page=" + strconv.Itoa(int(paginate.Page)-1)
	}

	if paginate.Page >= paginate.TotalPage {
		paginate.NextPage = ""
	} else {
		paginate.NextPage = viper.GetString("MAIN_URL") + "/api/" + viper.GetString("API_VERSION") + "/admin/articles?page=" + strconv.Itoa(int(paginate.Page)+1)
	}

	return articles, &paginate, nil
}

func (repository *ArticleRepositoryImpl) DeleteArticleById(id int) error {
	var article domain.Articles

	result := repository.db.Where("id = ?", id).Unscoped().Delete(&article)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Article not found")
	}
	return nil
}
