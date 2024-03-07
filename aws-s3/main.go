package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const region = "eu-central-1"
const uniqueBucketName = "bucket-125823"

type S3Client interface { // created for testing purposes
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
}

func main() {
	ctx := context.Background()

	s3Client, err := initS3Client(ctx)
	if err != nil {
		fmt.Printf("error while initializing s3 client, %v", err)
		os.Exit(1)
	}

	err2 := createS3Bucket(ctx, s3Client)
	if err2 != nil {
		fmt.Printf("error while creating s3 bucket, %v", err2)
		os.Exit(1)
	}

	err3 := uploadToS3Bucket(ctx, s3Client)
	if err3 != nil {
		fmt.Printf("error while uploading to s3 bucket, %v", err3)
		os.Exit(1)
	}

	fmt.Printf("Successfully uploaded file to s3!")
}

func initS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	return s3.NewFromConfig(cfg), nil
}

func createS3Bucket(ctx context.Context, s3Client S3Client) error {
	existingBuckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("error while listing buckets, %v", err)
	}

	for _, bucket := range existingBuckets.Buckets {
		if *bucket.Name == uniqueBucketName {
			return nil
		}
	}

	_, err2 := s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(uniqueBucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: region,
		},
	})
	if err2 != nil {
		return fmt.Errorf("unable to create s3 bucket, %v", err2)
	}

	return nil
}

func uploadToS3Bucket(ctx context.Context, s3Client *s3.Client) error {
	uploader := manager.NewUploader(s3Client)
	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(uniqueBucketName),
		Key:    aws.String("test.txt"),
		Body:   strings.NewReader("Hello, World!"),
	})
	if err != nil {
		return fmt.Errorf("error while uploading to s3 bucket, %v", err)
	}

	return nil
}
