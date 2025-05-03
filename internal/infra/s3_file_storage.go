package infra

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3FileStorage struct {
	client     *s3.Client
	bucketName string
	region     string
}

func NewS3FileStorage(bucketName, region string) (*S3FileStorage, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))

	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &S3FileStorage{
		client:     client,
		bucketName: bucketName,
		region:     region,
	}, nil
}

func (s *S3FileStorage) StoreFile(fileName, fileType string, fileContent []byte) (string, error) {
	_, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(fileContent),
		ContentType: aws.String(fileType),
		ACL:         types.ObjectCannedACLPublicRead,
	})

	if err != nil {
		return "", err
	}

	filePublicUrl := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucketName, s.region, fileName)

	return filePublicUrl, nil
}
