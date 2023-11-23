package initializers

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ConnectToS3(config *Config) (*s3.S3, error) {
	if config.S3BUCKET == "" {
		log.Fatal("Your S3 Bucket configuration was empty")
	}

	if config.S3ENDPOINT == "" {
		log.Fatal("Your S3 Endpoint configuration was empty")
	}

	if config.S3SECRETKEY == "" {
		log.Fatal("Your S3 Secret Key configuration was empty")
	}

	if config.S3ACCESSKEY == "" {
		log.Fatal("Your S3 Access Key configuration was empty")
	}

	if config.S3REGION == "" {
		log.Fatal("Your S3 Region configuration was empty")
	}

	// create an aws session
	s3session, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String(config.S3ENDPOINT),
		Region:      aws.String(config.S3REGION),
		Credentials: credentials.NewStaticCredentials(config.S3ACCESSKEY, config.S3SECRETKEY, ""),
	})

	if err != nil {
		return nil, err
	}

	// create a s3 client
	s3Client := s3.New(s3session)

	result, err := s3Client.ListBuckets(nil)
	if err != nil {
		return nil, err
	}

	log.Println("🚀 Connected Successfully to the AWS")
	for _, bucket := range result.Buckets {
		log.Println("list bucket " + *bucket.Name)
	}

	log.Println(config.S3BUCKET)
	_, err = s3Client.HeadBucket(&s3.HeadBucketInput{Bucket: aws.String(config.S3BUCKET)})

	if err != nil {
		log.Fatal("Bucket does not exist or could not be accessed:", err)
	}

	return s3Client, nil
}
