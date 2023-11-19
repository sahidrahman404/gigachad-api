package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	AWSBucket       string
	AWSRegion       string
}

func NewAWSConfig(cfg AWSConfig) (aws.Config, error) {
	return config.LoadDefaultConfig(
		context.Background(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AccessKeyID,
				cfg.SecretAccessKey,
				"",
			)),
		config.WithRegion(cfg.AWSRegion),
	)
}

func NewPresignClient(cfg aws.Config) *s3.PresignClient {
	client := s3.NewFromConfig(cfg)
	return s3.NewPresignClient(client)
}

func UploadURL(client *s3.PresignClient, fileKey string, cfg AWSConfig) (string, error) {
	expiration := time.Minute * 2
	putObjectArgs := s3.PutObjectInput{
		Bucket: &cfg.AWSBucket,
		Key:    &fileKey,
	}
	res, err := client.PresignPutObject(
		context.Background(),
		&putObjectArgs,
		s3.WithPresignExpires(expiration),
	)
	if err != nil {
		return "", err
	}
	return res.URL, nil
}
