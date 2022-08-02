package kinterfaceimpl

import (
	"github.com/aws/aws-sdk-go/service/secretsmanager"
<<<<<<< HEAD
=======
	"github.com/kraneware/kinterface/secrets"
>>>>>>> b24503ac92f974fca1167eba3d1c3a05b643ee39
	_ "github.com/kraneware/kinterface/secrets"
	"github.com/kraneware/kws/services"
)

<<<<<<< HEAD
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
=======
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
>>>>>>> b24503ac92f974fca1167eba3d1c3a05b643ee39
	if err != nil {
		return "", err
	}

	return data, nil
}

<<<<<<< HEAD
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
=======
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
>>>>>>> b24503ac92f974fca1167eba3d1c3a05b643ee39
}
