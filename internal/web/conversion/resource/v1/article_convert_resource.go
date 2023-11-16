package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/utils/helpers"
)

func ConvertArticleResource(articles []domain.Articles) []resources.ArticleResource {
	articleResource := []resources.ArticleResource{}
	for _, article := range articles {
		singleArticleResource := resources.ArticleResource{}
		singleArticleResource.Title = article.Title
		singleArticleResource.Thumbnail = article.Thumbnail
		singleArticleResource.Slug = article.Slug
		if article.Admin != nil {
			singleArticleResource.Author = resources.Author{
				Name: article.Admin.First_name + " " + article.Admin.Last_name,
				Role: article.Admin.Credential.Role.Name,
			}
		} else if article.Counselors != nil {
			singleArticleResource.Author = resources.Author{
				Name: article.Counselors.First_name + " " + article.Counselors.Last_name,
				Role: article.Counselors.Credential.Role.Name,
			}
		}
		singleArticleResource.Status = article.Status
		singleArticleResource.PublishedAt = helpers.ParseDateFormat(article.PublishedAt)

		articleResource = append(articleResource, singleArticleResource)
	}

	return articleResource
}

func ConvertSingleArticleResource(article *domain.Articles) resources.ArticleResource {
	singleArticleResource := resources.ArticleResource{}
	singleArticleResource.Id = article.Id
	singleArticleResource.Title = article.Title
	singleArticleResource.Thumbnail = article.Thumbnail
	singleArticleResource.Slug = article.Slug
	singleArticleResource.Content = article.Content
	if article.Admin != nil {
		singleArticleResource.Author = resources.Author{
			Name: article.Admin.First_name + " " + article.Admin.Last_name,
			Role: article.Admin.Credential.Role.Name,
		}
	} else if article.Counselors != nil {
		singleArticleResource.Author = resources.Author{
			Name: article.Counselors.First_name + " " + article.Counselors.Last_name,
			Role: article.Counselors.Credential.Role.Name,
		}
	}
	singleArticleResource.Status = article.Status
	singleArticleResource.PublishedAt = helpers.ParseDateFormat(article.PublishedAt)

	return singleArticleResource
}
