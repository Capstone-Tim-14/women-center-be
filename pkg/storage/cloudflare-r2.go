package storage

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
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
		config.WithRegion("apac"),
	)

	if ErrConfiguration != nil {
		panic(ErrConfiguration.Error())
	}

	client := s3.NewFromConfig(configuration)

	return client

}

func S3PutFile(file *multipart.FileHeader, directory string) (string, error) {

	bucket := viper.GetString("STORAGE.BUCKET_NAME")
	S3 := S3Config()

	src, err := file.Open()

	if err != nil {
		return "", err
	}

	defer src.Close()

	generateKey := fmt.Sprintf("%s/%s", directory, uuid.NewV4().String())

	uploadInput := &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &generateKey,
		Body:        src,
		ContentType: aws.String("image/png"),
	}

	_, errObject := S3.PutObject(context.TODO(), uploadInput)

	if errObject != nil {
		return "", errObject
	}

	publicURL := fmt.Sprintf("https://pub-31b2a1a1e015474f97220ee42fe1d856.r2.dev/%s", generateKey)

	return publicURL, nil

}
