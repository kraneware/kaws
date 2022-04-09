package kinterfaceimpl

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kraneware/kaws"
	"github.com/kraneware/kaws/ks3"
	_ "github.com/kraneware/kws/services"
)

type S3Persister struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	client ks3.Client

	BucketName     string `json:"bucket"`
	BucketRootPath string `json:"bucket_root_path"`
}

func CreatePersister(name string, id string) (S3Persister, error) {
	p := S3Persister{
		Id:   id,
		Name: name,
	}

	return p, nil
}

func (t S3Persister) ListBuckets() ([]s3.Bucket, error) {
	buckets, err := t.client.GetBucketList()
	return buckets, err
}

func (t S3Persister) SetId(id string) {
	t.Id = id
}

func (t S3Persister) LoadFile(path string, name string) (content string) {
	return ""
}

func (t S3Persister) SaveFile(path string, name string, content string) (string, string, error) {
	return "", "", nil
}

func (t S3Persister) InitPersister() (S3Persister, error) {
	var err error

	t.client, err = ks3.CreateClient()
	return t, err
}

func (t S3Persister) ClosePersister() error {
	var innerClient *s3.S3 = t.client.Client()

	if kaws.DEBUG {
		_ = fmt.Sprintf("%v", innerClient)
	}

	return nil
}

func (t S3Persister) GetConnection() (interface{}, error) {
	if &t != nil && &t.client != nil {
		return t.client.Client(), nil
	} else {
		return nil, errors.New("sender is nil")
	}
}

func (t S3Persister) ID() string {
	return t.Id
}

func (t S3Persister) GetName() string {
	return t.Name
}

func (t S3Persister) SetName(name string) {
	t.Name = name
}
