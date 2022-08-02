package kinterfaceimpl

import (
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	_ "github.com/kraneware/kinterface/secrets"
	"github.com/kraneware/kws/services"
)

type SecretsKeeper struct {
	Id     string                        `json:"id"`
	Name   string                        `json:"name"`
	Client secretsmanager.SecretsManager `json:"client"`
}

func (t *SecretsKeeper) SetID(id string) {
	t.Id = id
}

func (t *SecretsKeeper) ID() string {
	return t.Id
}

func (t *SecretsKeeper) GetName() string {
	return t.Name
}

func (t *SecretsKeeper) SetName(name string) {
	t.Name = name
}

func (t *SecretsKeeper) LoadSecret(path string) (string, error) {
	data, err := services.GetSecretByArn(&t.Client, path)
	if err != nil {
		return "", err
	}

	return data, nil
}

func (t *SecretsKeeper) SaveSecret(path string, content string) error {
	return nil
}

func (t *SecretsKeeper) InitSecretsKeeper() error {
	t.Client = *services.SecretClient()
	return nil
}

func (t *SecretsKeeper) CloseSecretsKeeper() error {
	return nil
}

func (t *SecretsKeeper) GetConnection() (interface{}, error) {
	return t.Client, nil
}
