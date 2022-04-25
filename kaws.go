package kaws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kraneware/kws/services"
	"os"
)

const (
	DEBUG = false
	Name  = "kaws"
	Id    = "kaws"
)

func AwsConnectionValid() error {
	s3c := services.S3Client()
	out, err := s3c.ListBuckets(&s3.ListBucketsInput{})

	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "%s", err.Error())
		return err
	}
	if len(out.Buckets) <= 0 {
		_, _ = fmt.Fprintf(os.Stdout, "%s \n", "aws connection valid, but no buckets found")
		return nil
	}

	return nil
}
