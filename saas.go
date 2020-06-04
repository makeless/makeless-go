package go_saas

import (
	"fmt"
	"sync"

	"github.com/go-saas/go-saas/api"
	"github.com/go-saas/go-saas/database"
	"github.com/go-saas/go-saas/logger"
)

type Saas struct {
	License  string
	Logger   go_saas_logger.Logger
	Database *saas_database.Database
	Api      *saas_api.Api
	*sync.RWMutex
}

func (saas Saas) getLicense() string {
	saas.RLock()
	defer saas.RUnlock()

	return saas.License
}

func (saas Saas) isLicenseValid() bool {
	saas.getLicense()
	return true
}

func (saas Saas) GetLogger() go_saas_logger.Logger {
	saas.RLock()
	defer saas.RUnlock()

	return saas.Logger
}

func (saas Saas) GetDatabase() *saas_database.Database {
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
	if !saas.isLicenseValid() {
		return fmt.Errorf("invalid saas license")
	}

	if err := saas.GetDatabase().Connect(); err != nil {
		return err
	}

	if err := saas.GetDatabase().AutoMigrate(); err != nil {
		return err
	}

	if err := saas.GetApi().Start(); err != nil {
		return err
	}

	return nil
}
