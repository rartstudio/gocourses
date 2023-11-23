package clients

import (
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rartstudio/gocourses/initializers"
)

func (s S3Service) UploadFileToS3(file *multipart.File, fileInfo *multipart.FileHeader) (string, error) {
	// file extension
	fileExtension := filepath.Ext(fileInfo.Filename)

	// filepath
	objectKey := "uploads/" + strconv.Itoa(int(time.Now().Unix())) + fileExtension

	data := &s3.PutObjectInput{
		Bucket: aws.String(s.config.S3BUCKET),
		Key:    aws.String(objectKey),
		Body:   *file,
	}

	_, err := s.s3Client.PutObject(data)

	return objectKey, err
}

func (s S3Service) RemoveFileFromS3(key string) error {
	data := &s3.DeleteObjectInput{
		Bucket: aws.String(s.config.S3BUCKET),
		Key:    aws.String(key),
	}
	_, err := s.s3Client.DeleteObject(data)

	return err
}

type S3Service struct {
	s3Client *s3.S3
	config   *initializers.Config
}

func NewS3Service(s3Client *s3.S3, config *initializers.Config) *S3Service {
	return &S3Service{
		s3Client: s3Client,
		config:   config,
	}
}
