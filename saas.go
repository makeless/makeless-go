package go_saas

import (
	"sync"

	"github.com/go-saas/go-saas/database"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/logger"
)

type Saas struct {
	Logger   go_saas_logger.Logger
	Database go_saas_database.Database
	Http     go_saas_http.Http
	*sync.RWMutex
}

func (saas *Saas) GetLogger() go_saas_logger.Logger {
	saas.RLock()
	defer saas.RUnlock()

	return saas.Logger
}

func (saas *Saas) GetDatabase() go_saas_database.Database {
	saas.RLock()
	defer saas.RUnlock()

	return saas.Database
}

func (saas *Saas) GetHttp() go_saas_http.Http {
	saas.RLock()
	defer saas.RUnlock()

	return saas.Http
}

func (saas *Saas) SetRoute(name string, handler func(http go_saas_http.Http) error) {
	saas.GetHttp().SetHandler(name, handler)
}

func (saas *Saas) Init() error {
	if err := saas.GetDatabase().Connect(); err != nil {
		return err
	}

	if err := saas.GetDatabase().Migrate(); err != nil {
		return err
	}

	saas.SetRoute("ok", saas.ok)

	return nil
}

func (saas *Saas) Run() error {
	return saas.GetHttp().Start()
}
