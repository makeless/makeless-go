package go_saas

import (
	"fmt"
	"github.com/loeffel-io/go-saas/api"
	"github.com/loeffel-io/go-saas/database"
	"github.com/loeffel-io/go-saas/logger"
	"sync"
)

type Saas struct {
	License  string
	Logger   saas_logger.Logger
	Database *saas_database.Database
	Api      *saas_api.Api
	*sync.RWMutex
}

func (saas Saas) getLicense() string {
	saas.RLock()
	defer saas.RUnlock()

	return saas.License
}

func (saas Saas) validateLicense() bool {
	saas.getLicense()
	return true
}

func (saas Saas) GetLogger() saas_logger.Logger {
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
	if !saas.validateLicense() {
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
