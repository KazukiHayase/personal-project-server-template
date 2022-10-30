package resolver

import (
	"errors"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/KazukiHayase/server-template/config"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Config    config.Config
	Datastore datastore.Client
}

func NewResolver(config config.Config, datastore datastore.Client) Resolver {
	return Resolver{
		Config:    config,
		Datastore: datastore,
	}
}

func (r *Resolver) NewSystemError(e error) error {
	log.Println("System Error: ", e.Error())
	return errors.New("システムエラーが発生しました")
}

func (r *Resolver) NewBusinessError(e error) error {
	log.Println("Business Error: ", e.Error())
	return e
}
