package ks3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kraneware/kaws"
	"github.com/kraneware/kws/services"
)

type Client struct {
	client *s3.S3
}

func CreateClient() (Client, error) {
	s3Client := services.S3Client()

	c := Client{
		client: s3Client,
	}

	_, _ = c.GetBucketList()

	return c, nil
}

func (t Client) GetBucketObjects(bucketName string, key string) ([]s3.Object, error) {
	lib := &s3.ListObjectsV2Input{
		Bucket:              &bucketName,
		Delimiter:           nil,
		EncodingType:        nil,
		ExpectedBucketOwner: nil,
		MaxKeys:             nil,
		Prefix:              &key,
		RequestPayer:        nil,
	}

	out, err := t.client.ListObjectsV2(lib)
	if err != nil {
		return nil, err
	}

	var objects []s3.Object
	for _, b := range out.Contents {
		if kaws.DEBUG {
			_ = fmt.Sprintf("object: %s \n", b.String())
		}

		objects = append(objects, *b)
	}

	return objects, err
}

func (t Client) GetBucketList() ([]s3.Bucket, error) {
	lib := &s3.ListBucketsInput{}

	out, err := t.client.ListBuckets(lib)
	if err != nil {
		return nil, err
	}

	var buckets []s3.Bucket
	for _, b := range out.Buckets {
		if kaws.DEBUG {
			_ = fmt.Sprintf("bucket: %s \n", *b.Name)
		}

		buckets = append(buckets, *b)
	}

	return buckets, err
}

func (t Client) Client() *s3.S3 {
	return t.client
}

func (t Client) SetClient(client *s3.S3) {
	t.client = client
}
