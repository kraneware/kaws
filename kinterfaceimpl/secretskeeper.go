package kinterfaceimpl

import (
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/kraneware/kinterface/secrets"
	_ "github.com/kraneware/kinterface/secrets"
	"github.com/kraneware/kws/services"
)

type Skeeper struct {
	id     string                        `json:"id"`
	name   string                        `json:"name"`
	client secretsmanager.SecretsManager `json:"client"`
}

func (Skeeper) ID() string {
	//TODO implement me
	panic("implement me")
}

func (Skeeper) SetId(id string) {
	//TODO implement me
	panic("implement me")
}

func (Skeeper) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (Skeeper) SetName(name string) {
	//TODO implement me
	panic("implement me")
}

func (t Skeeper) LoadSecret(path string) (string, error) {
	data, err := services.GetSecretByArn(&t.client, path)
	if err != nil {
		return "", err
	}

	return data, nil
}

func (t Skeeper) SaveSecret(path string, content string) error {
	t.client = *services.SecretClient()
	return nil
}

func (t Skeeper) InitSecretsKeeper() error {
	t.client = *services.SecretClient()
	return nil
}

func (Skeeper) CloseSecretsKeeper() error {
	//TODO implement me
	panic("implement me")
}

func (Skeeper) GetConnection() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func CreateKeeper() (secrets.Keeper, error) {
	sk := Skeeper{}
	return sk, nil
}
