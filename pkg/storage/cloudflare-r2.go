package storage

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

const (
	article_dir = "/articles/thumbnail/"
)

func S3Config() *s3.Client {

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
				URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", viper.GetString("STORAGE.ACCOUNT_ID")),
			},
			nil
	})

	configuration, ErrConfiguration := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(viper.GetString("STORAGE.ACCESS_KEY_ID"), viper.GetString("STORAGE.SECRET_ACCESS_KEY"), "")),
	)

	if ErrConfiguration != nil {
		panic(ErrConfiguration.Error())
	}

	client := s3.NewFromConfig(configuration)

	return client

}

func UploadFileImage(file *multipart.FileHeader, category string) (string, error) {

	bucket := viper.GetString("STORAGE.BUCKET_NAME")
	S3 := S3Config()

	src, err := file.Open()

	if err != nil {
		return "", err
	}

	defer src.Close()

	uploadInput := &s3.PutObjectInput{
		Bucket:      &bucket,
		Body:        src,
		ContentType: aws.String("image/png"),
	}

	presignClient := s3.NewPresignClient(S3)

	presignResp, errObject := presignClient.PresignPutObject(context.TODO(), uploadInput)

	if errObject != nil {
		return "", errObject
	}

	return presignResp.URL, nil

}
