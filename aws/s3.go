package aws

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Info struct {
	region     string
	bucketName string
	conf       *aws.Config
}

func NewS3Info(region, bucketName string) *S3Info {
	conf := &aws.Config{
		Credentials: credentials.NewEnvCredentials(),
		Region:      aws.String(region),
	}

	return &S3Info{
		region:     region,
		bucketName: bucketName,
		conf:       conf,
	}
}

func (info *S3Info) Upload(f *os.File, key string) (string, error) {
	uploader := s3manager.NewUploader(session.New(info.conf))
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   f,
		Bucket: aws.String(info.bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Printf("Failed to upload: %s", err)
		return "", err
	}

	return result.Location, nil
}
