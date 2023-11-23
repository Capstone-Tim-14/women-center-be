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
	FindAllArticle(string, string, query.Pagination) ([]domain.Articles, *query.Pagination, error)
	FindById(id int) (*domain.Articles, error)
	DeleteArticleById(id int) error
	UpdateStatusArticle(slug, status string) error
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

func (repository *ArticleRepositoryImpl) FindAllArticle(orderBy string, search string, paginate query.Pagination) ([]domain.Articles, *query.Pagination, error) {
	var articles []domain.Articles

	result := repository.db.Scopes(query.Paginate(articles, &paginate, repository.db)).Preload("Admin").Preload("Admin.Credential").Preload("Admin.Credential.Role").Preload("Counselors").Preload("Counselors.Credential").Preload("Counselors.Credential.Role").Where("title LIKE ?", "%"+search+"%")

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

func (repository *ArticleRepositoryImpl) FindById(id int) (*domain.Articles, error) {
	var article domain.Articles

	result := repository.db.Where("id = ?", id).First(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Article not found")
	}
	return &article, nil
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

func (repository *ArticleRepositoryImpl) UpdateStatusArticle(slug, status string) error {

	result := repository.db.Table("articles").Where("slug = ?", slug).Updates(domain.Articles{Status: status})
	if result.Error != nil {
		return fmt.Errorf("Error update status article")
	}

	return nil
}

func (repository *ArticleRepositoryImpl) FindBySlug(slug string) (*domain.Articles, error) {
	article := domain.Articles{}
	result := repository.db.Preload("Admin").Preload("Admin.Credential").Preload("Admin.Credential.Role").Preload("Counselors").Preload("Counselors.Credential").Preload("Counselors.Credential.Role").Preload("Tags").Where("slug = ?", slug).First(&article)
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
