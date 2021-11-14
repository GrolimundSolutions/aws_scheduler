package utils

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func Test(DbId string) {
	// Load the Shared AWS Configuration (~/.aws/config)
	// or load from the environment variables
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	s3Client := s3.NewFromConfig(cfg)
	rdsClient := rds.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("my-bucket"),
	})
	if err != nil {
		log.Fatal(err)
	}

	rdsOutput, err := rdsClient.StopDBInstance(context.TODO(), &rds.StopDBInstanceInput{
		DBInstanceIdentifier: &DbId,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}

	log.Println("rdsOutput results:", rdsOutput)

}
