package application

import (
	"app/core/server"
	"app/pkg/bootstrap"
)

type Application struct {
	Config     bootstrap.Config
	Server     server.Server
	Singletons bootstrap.Singletons
}

func Create() (Application, error) {
	config, err := bootstrap.InitConfig()
	if err != nil {
		return Application{}, err
	}

	singletons, err := bootstrap.InitSingletons(config)
	if err != nil {
		return Application{}, err
	}

	app := Application{
		Config:     config,
		Server:     server.New(config.Server),
		Singletons: singletons,
	}

	return app, nil
}
