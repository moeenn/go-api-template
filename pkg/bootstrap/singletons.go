package bootstrap

import (
	"app/core/database"
)

type Singletons struct {
	Database database.Database
}

func InitSingletons(c Config) (Singletons, error) {
	var singletons Singletons

	if err := singletons.Database.Connect(c.Database); err != nil {
		return Singletons{}, err
	}

	return singletons, nil
}
