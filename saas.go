package go_saas

import (
	"github.com/go-saas/go-saas/database"
	"sync"

	"github.com/go-saas/go-saas/api"
	"github.com/go-saas/go-saas/logger"
)

type Saas struct {
	Logger   go_saas_logger.Logger
	Database go_saas_database.Database
	Api      *saas_api.Api
	*sync.RWMutex
}

func (saas Saas) GetLogger() go_saas_logger.Logger {
	saas.RLock()
	defer saas.RUnlock()

	return saas.Logger
}

func (saas Saas) GetDatabase() go_saas_database.Database {
	saas.RLock()
	defer saas.RUnlock()

	return saas.Database
}

func (saas Saas) GetApi() *saas_api.Api {
	saas.RLock()
	defer saas.RUnlock()

	return saas.Api
}

func (saas Saas) Run() error {
	if err := saas.GetDatabase().Connect(); err != nil {
		return err
	}

	if err := saas.GetDatabase().Migrate(); err != nil {
		return err
	}

	if err := saas.GetApi().Start(); err != nil {
		return err
	}

	return nil
}
