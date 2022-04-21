package kinterfaceimpl

import (
	_ "github.com/kraneware/kinterface/secrets"
	"github.com/kraneware/kws/services"
)
import "github.com/aws/aws-sdk-go/service/secretsmanager"

type SecretsKeeper struct {
	id     string                        `json:"id"`
	name   string                        `json:"name"`
	client secretsmanager.SecretsManager `json:"client"`
}

func (t *SecretsKeeper) SetID(id string) {
	t.id = id
}

func (t *SecretsKeeper) ID() string {
	return t.id
}

func (t *SecretsKeeper) GetName() string {
	return t.name
}

func (t *SecretsKeeper) SetName(name string) {
	t.name = name
}

func (t *SecretsKeeper) LoadSecret(path string) (string, error) {
	data, err := services.GetSecretByArn(&t.client, path)
	if err != nil {
		return "", err
	}

	return data, nil
}

func (t *SecretsKeeper) SaveSecret(path string, content string) error {

}

func (t *SecretsKeeper) InitSecretsKeeper() error {
	t.client = *services.SecretClient()
	return nil
}

func (t *SecretsKeeper) CloseSecretsKeeper() error {

}

func (t *SecretsKeeper) GetConnection() (interface{}, error) {
	return t.client, nil
}
