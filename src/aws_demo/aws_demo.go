package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	fmt.Println("AWS S3 Demo")

	awsConfig := aws.Config{Region: aws.String("us-west-2")}

	s3 := s3.New(&awsConfig)

	resp, err := s3.ListBuckets(nil)

	fmt.Println(*s3, resp, err)

	fmt.Println("Done.")
}
