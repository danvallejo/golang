package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	fmt.Println("AWS S3 Demo")

	awsConfig := aws.Config{Region: aws.String("us-west-2")}

	s3 := s3.New(&awsConfig)

	resp, err := s3.ListBuckets(nil)

	fmt.Println(*s3, resp, err)
	
	db := dynamodb.New(&awsConfig)

	dbTables, err := db.ListTables(nil)
	fmt.Println(*db, dbTables, err)

	fmt.Println("Done.")
}
