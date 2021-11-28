package kaws

import (
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/kraneware/core-go/awsutil/services"
	"github.com/kraneware/kinterface/persistance"
)

type Ks3 struct {
	Id					string 		`json:"id"`
	Name				string		`json:"name"`

	s3Client			s3.S3

	BucketName			string		`json:"bucket"`
	BucketRootPath		string		`json:"bucket_root_path"`
}

func (t Ks3) SetId(id string) {
	t.Id = id
}

func (t Ks3) LoadFile(path string, name string) (content string) {
	return ""
}

func (t Ks3) SaveFile(path string, name string, content string) (string, string, error) {
	return "", "", nil
}

func (t Ks3) CreatePersister(name string, id string) (persistance.Persister, error) {
	p := Ks3{
		Id:   id,
		Name: name,
	}

	return p, nil
}

func (t Ks3) InitPersister() (persistance.Persister, error) {
	return &t, nil
}

func (t Ks3) ClosePersister() error {
	return nil
}

func (t Ks3) GetPersister() (interface{}, error) {
	return t.s3Client, nil
}

func (t Ks3) ID() string {
	return t.Id
}

func (t Ks3) SetID(id string) {
	t.Id = id
}

func (t Ks3) GetName() string {
	return t.Name
}

func (t Ks3) SetName(name string) {
	t.Name = name
}