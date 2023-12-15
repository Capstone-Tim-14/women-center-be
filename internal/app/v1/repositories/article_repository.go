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
	GetLatestArticleForUser() (*domain.Articles, error)
	GetListArticleForUser() ([]domain.Articles, error)
	FindAllArticle(string, string, query.Pagination) ([]domain.Articles, *query.Pagination, error)
	FindArticleCounselor(id int) ([]domain.Articles, *domain.ArticleStatusCount, error)
	FindById(id int) (*domain.Articles, error)
	DeleteArticleById(id int) error
	UpdateStatusArticle(slug, status string) error
	FindActiveArticleBySlug(slug string) (*domain.Articles, error)
	FindBySlug(slug string) (*domain.Articles, error)
	FindByTitle(title string) (*domain.Articles, error)
	UpdateArticle(id int, article *domain.Articles) error
	FindSlugForFavorite(slug string) (*domain.Articles, error)
}

type ArticleRepositoryImpl struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &ArticleRepositoryImpl{
		db: db,
	}
}

func (repository *ArticleRepositoryImpl) FindActiveArticleBySlug(slug string) (*domain.Articles, error) {
	article := domain.Articles{}
	result := repository.db.Preload("Admin").Preload("Admin.Credential").Preload("Admin.Credential.Role").Preload("Counselors").Preload("Counselors.Credential").Preload("Counselors.Credential.Role").Preload("Tags").Where("slug = ? AND status = ?", slug, "PUBLISHED").First(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return &article, nil
}

func (repository *ArticleRepositoryImpl) FindArticleCounselor(id int) ([]domain.Articles, *domain.ArticleStatusCount, error) {

	var articles []domain.Articles
	var articles_count *domain.ArticleStatusCount

	errListArticles := repository.db.InnerJoins("Counselors").InnerJoins("Counselors.Credential").InnerJoins("Counselors.Credential.Role").Find(&articles, "Counselors.id = ?", id)

	if errListArticles.Error != nil {
		fmt.Println(errListArticles.Error.Error())
		return nil, nil, fmt.Errorf("Error to find article")
	}

	if errListArticles.RowsAffected == 0 {
		return nil, nil, fmt.Errorf("Articles empty")
	}

	errGetCountArticles := repository.db.Raw("SELECT SUM(CASE WHEN articles.status = 'PUBLISHED' THEN 1 ELSE 0 END) AS PUBLISHED_COUNT, SUM(CASE WHEN articles.status = 'REVIEW' THEN 1 ELSE 0 END) AS REVIEW_COUNT, SUM(CASE WHEN articles.status = 'REJECTED' THEN 1 ELSE 0 END) AS REJECTED_COUNT FROM `articles` INNER JOIN `counselors` ON counselors.id = articles.counselors_id WHERE counselors.id = ?", id).Scan(&articles_count)

	if errGetCountArticles.Error != nil {
		return nil, nil, fmt.Errorf("Error to generate count article")
	}

	return articles, articles_count, nil

}

func (repository *ArticleRepositoryImpl) CreateArticle(article *domain.Articles) (*domain.Articles, error) {
	result := repository.db.Create(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return article, nil

}

func (repository *ArticleRepositoryImpl) GetLatestArticleForUser() (*domain.Articles, error) {

	var article *domain.Articles

	errTakeArticle := repository.db.Preload("Admin").Preload("Admin.Credential").Preload("Admin.Credential.Role").Preload("Counselors").Preload("Counselors.Credential").Preload("Counselors.Credential.Role").Order("published_at desc").Take(&article, "status = ?", "PUBLISHED")

	if errTakeArticle.Error != nil {
		fmt.Errorf(errTakeArticle.Error.Error())
		return nil, fmt.Errorf("Error to find article")
	}

	if errTakeArticle.RowsAffected == 0 {
		return nil, fmt.Errorf("Article not found")
	}

	return article, nil
}

func (repository *ArticleRepositoryImpl) GetListArticleForUser() ([]domain.Articles, error) {

	var articles []domain.Articles
	errListArticles := repository.db.Preload("Admin").Preload("Admin.Credential").Preload("Admin.Credential.Role").Preload("Counselors").Preload("Counselors.Credential").Preload("Counselors.Credential.Role").Find(&articles, "status = ?", "PUBLISHED")

	if errListArticles.Error != nil {
		fmt.Errorf(errListArticles.Error.Error())
		return nil, fmt.Errorf("Error to show articles list")
	}

	if errListArticles.RowsAffected == 0 {
		return nil, fmt.Errorf("Articles is empty")
	}

	return articles, nil

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

	result := repository.db.Preload("Tags").Where("id = ?", id).First(&article)
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

	transaction := repository.db.Begin()

	result := transaction.Where("id = ?", id).First(&article)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Article not found")
	}

	if errAssociation := transaction.Model(&article).Association("Tags").Clear(); errAssociation != nil {
		transaction.Rollback()
		fmt.Errorf(errAssociation.Error())
		return fmt.Errorf("Error when delete relations")
	}

	if errDeleteArticle := transaction.Unscoped().Delete(&article).Error; errDeleteArticle != nil {
		transaction.Rollback()
		fmt.Errorf(errDeleteArticle.Error())
		return fmt.Errorf("Error when delete article")
	}

	transaction.Commit()

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

func (repository *ArticleRepositoryImpl) UpdateArticle(id int, article *domain.Articles) error {
	result := repository.db.Model(&article).Where("id = ?", id).Updates(article)
	if result.Error != nil {
		return fmt.Errorf("error updating article")
	}

	return nil
}

func (repository *ArticleRepositoryImpl) FindSlugForFavorite(slug string) (*domain.Articles, error) {
	article := domain.Articles{}
	result := repository.db.Where("slug = ?", slug).First(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return &article, nil
}
